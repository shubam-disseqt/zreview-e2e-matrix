// SPDX-License-Identifier: Apache-2.0
// Copyright 2026 shubam-ai-code-reviewer contributors

// assert compares the JSON output of `sacr review --format json` against
// a case's expected findings and reports which seeded bugs were caught.
//
// Usage:
//
//	assert -expected .matrix/expected.json -actual out.json [-case name] [-json]
//
// Match rules — an actual finding matches an expected finding when:
//  1. category is equal (case-insensitive)
//  2. path is equal
//  3. actual line range overlaps the expected line range, allowing
//     LINE_TOLERANCE lines of slack on either side (LLMs off-by-N)
//  4. actual severity is >= expected minimum severity
//
// Extra actual findings (no match in expected) are reported but do not
// affect the match count — precision is tracked separately and does not
// gate CI in this phase.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LineTolerance is the number of lines an actual finding may drift from the
// expected range and still count as a match. LLMs frequently emit line
// numbers +/- 1-2 lines from the actual anchor; 3 is empirically generous.
const LineTolerance = 3

var severityOrder = map[string]int{
	"low":      1,
	"medium":   2,
	"high":     3,
	"critical": 4,
}

// ExpectedFile is the ground-truth schema stored per case at
// .matrix/expected.json on that case's branch.
type ExpectedFile struct {
	Case        string     `json:"case"`
	Description string     `json:"description,omitempty"`
	Findings    []Expected `json:"findings"`
}

// Expected is one seeded bug the reviewer must catch.
type Expected struct {
	Path        string `json:"path"`
	StartLine   int    `json:"start_line"`
	EndLine     int    `json:"end_line"`
	Category    string `json:"category"`     // bug | security | performance | maintainability | test | style | documentation | other
	MinSeverity string `json:"min_severity"` // low | medium | high | critical
	Description string `json:"description,omitempty"`
	// Optional. When true, missing this finding does NOT fail the case;
	// it is reported as a soft miss.
	Soft bool `json:"soft,omitempty"`
}

// ActualFile mirrors the shape emitted by `sacr review --format json`.
type ActualFile struct {
	Comments []Actual `json:"comments"`
}

type Actual struct {
	Path      string `json:"path"`
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	Category  string `json:"category"`
	Severity  string `json:"severity"`
	Content   string `json:"content"`
}

// Result is the assert program's structured output. With several -actual
// files, Matched/MatchedHard count expected findings caught in EVERY run
// (intersection) and drive the gate; MatchedUnion counts those caught in
// at least one run.
type Result struct {
	Case             string   `json:"case"`
	Description      string   `json:"description,omitempty"`
	Runs             int      `json:"runs"`
	ExpectedTotal    int      `json:"expected_total"`
	ExpectedHard     int      `json:"expected_hard"`
	Matched          int      `json:"matched"`
	MatchedHard      int      `json:"matched_hard"`
	MatchedUnion     int      `json:"matched_union"`
	MatchedUnionHard int      `json:"matched_union_hard"`
	UnionRecall      float64  `json:"union_recall"` // hard, any run
	SoftMisses       []string `json:"soft_misses,omitempty"`
	HardMisses       []string `json:"hard_misses,omitempty"`
	ExtraFindings    int      `json:"extra_findings"`
	Recall           float64  `json:"recall"`
	HardRecall       float64  `json:"hard_recall"`
	Passed           bool     `json:"passed"`
	MissedCritical   bool     `json:"missed_critical"`
	CriticalSummary  string   `json:"critical_summary,omitempty"`
}

// Match reports whether actual satisfies expected under the rules above.
func Match(exp Expected, act Actual) bool {
	if !strings.EqualFold(exp.Category, act.Category) {
		return false
	}
	if exp.Path != act.Path {
		return false
	}
	need, ok := severityOrder[strings.ToLower(exp.MinSeverity)]
	if !ok {
		need = severityOrder["low"]
	}
	got := severityOrder[strings.ToLower(act.Severity)]
	if got < need {
		return false
	}
	expStart := exp.StartLine - LineTolerance
	expEnd := exp.EndLine + LineTolerance
	if expStart < 1 {
		expStart = 1
	}
	if act.EndLine < act.StartLine {
		act.EndLine = act.StartLine
	}
	return maxInt(expStart, act.StartLine) <= minInt(expEnd, act.EndLine)
}

// matchRun reports, per expected finding, whether one run caught it, and
// how many of the run's comments matched nothing (extras).
func matchRun(exp ExpectedFile, act ActualFile) (hits []bool, extras int) {
	hits = make([]bool, len(exp.Findings))
	usedActual := make([]bool, len(act.Comments))
	for ei, e := range exp.Findings {
		for i, a := range act.Comments {
			if usedActual[i] {
				continue
			}
			if Match(e, a) {
				usedActual[i] = true
				hits[ei] = true
				break
			}
		}
	}
	for _, used := range usedActual {
		if !used {
			extras++
		}
	}
	return hits, extras
}

// Assert consumes one expected file and one or more actual runs. A finding
// counts as matched only when every run caught it; union is reported too.
func Assert(caseName string, exp ExpectedFile, runs ...ActualFile) Result {
	r := Result{
		Case:          caseName,
		Description:   exp.Description,
		Runs:          len(runs),
		ExpectedTotal: len(exp.Findings),
	}
	inAll := make([]bool, len(exp.Findings))
	inAny := make([]bool, len(exp.Findings))
	for i := range inAll {
		inAll[i] = len(runs) > 0
	}
	for _, act := range runs {
		hits, extras := matchRun(exp, act)
		r.ExtraFindings += extras
		for i, h := range hits {
			inAll[i] = inAll[i] && h
			inAny[i] = inAny[i] || h
		}
	}
	for ei, e := range exp.Findings {
		if !e.Soft {
			r.ExpectedHard++
		}
		if inAny[ei] {
			r.MatchedUnion++
			if !e.Soft {
				r.MatchedUnionHard++
			}
		}
		if inAll[ei] {
			r.Matched++
			if !e.Soft {
				r.MatchedHard++
			}
			continue
		}
		miss := fmt.Sprintf("%s:%d-%d [%s/%s]%s",
			e.Path, e.StartLine, e.EndLine, e.Category, e.MinSeverity, describe(e))
		if e.Soft {
			r.SoftMisses = append(r.SoftMisses, miss)
		} else {
			r.HardMisses = append(r.HardMisses, miss)
		}
		if !e.Soft && strings.EqualFold(e.MinSeverity, "critical") {
			r.MissedCritical = true
			if r.CriticalSummary != "" {
				r.CriticalSummary += "; "
			}
			r.CriticalSummary += miss
		}
	}
	if r.ExpectedTotal > 0 {
		r.Recall = float64(r.Matched) / float64(r.ExpectedTotal)
	}
	if r.ExpectedHard > 0 {
		r.HardRecall = float64(r.MatchedHard) / float64(r.ExpectedHard)
		r.UnionRecall = float64(r.MatchedUnionHard) / float64(r.ExpectedHard)
	} else {
		r.HardRecall = 1.0
		r.UnionRecall = 1.0
	}
	r.Passed = r.ExpectedHard == r.MatchedHard && !r.MissedCritical
	return r
}

func describe(e Expected) string {
	if e.Description == "" {
		return ""
	}
	return " — " + e.Description
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// --- CLI ---

func main() {
	expectedPath := flag.String("expected", ".matrix/expected.json", "path to expected.json")
	var actualPaths []string
	flag.Func("actual", "path to sacr review json output (repeat for multiple runs; the gate uses their intersection)", func(v string) error {
		actualPaths = append(actualPaths, v)
		return nil
	})
	caseName := flag.String("case", "", "case name (defaults to the expected.json 'case' field or the parent dir)")
	asJSON := flag.Bool("json", false, "emit machine-readable JSON instead of markdown")
	flag.Parse()

	exp, err := readExpected(*expectedPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "assert: read expected: %v\n", err)
		os.Exit(2)
	}
	if len(actualPaths) == 0 {
		actualPaths = []string{"out.json"}
	}
	runs := make([]ActualFile, 0, len(actualPaths))
	for _, ap := range actualPaths {
		act, err := readActual(ap)
		if err != nil {
			fmt.Fprintf(os.Stderr, "assert: read actual %s: %v\n", ap, err)
			os.Exit(2)
		}
		runs = append(runs, act)
	}
	name := *caseName
	if name == "" {
		name = exp.Case
	}
	if name == "" {
		name = filepath.Base(filepath.Dir(*expectedPath))
	}

	r := Assert(name, exp, runs...)

	if *asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(r)
	} else {
		writeMarkdown(os.Stdout, r)
	}
	if !r.Passed {
		os.Exit(1)
	}
}

func readExpected(path string) (ExpectedFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ExpectedFile{}, err
	}
	var e ExpectedFile
	if err := json.Unmarshal(data, &e); err != nil {
		return ExpectedFile{}, fmt.Errorf("parse %s: %w", path, err)
	}
	return e, nil
}

func readActual(path string) (ActualFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ActualFile{}, err
	}
	var a ActualFile
	if err := json.Unmarshal(data, &a); err != nil {
		return ActualFile{}, fmt.Errorf("parse %s: %w", path, err)
	}
	return a, nil
}

func writeMarkdown(w *os.File, r Result) {
	status := ":white_check_mark: pass"
	if !r.Passed {
		status = ":x: fail"
	}
	fmt.Fprintf(w, "### %s — %s\n\n", r.Case, status)
	if r.Description != "" {
		fmt.Fprintf(w, "%s\n\n", r.Description)
	}
	fmt.Fprintf(w, "- runs: %d\n", r.Runs)
	fmt.Fprintf(w, "- expected (hard): %d\n", r.ExpectedHard)
	fmt.Fprintf(w, "- matched (hard, every run): %d\n", r.MatchedHard)
	fmt.Fprintf(w, "- hard recall (intersection): %.0f%%\n", r.HardRecall*100)
	fmt.Fprintf(w, "- union recall (any run): %.0f%%\n", r.UnionRecall*100)
	fmt.Fprintf(w, "- extra findings: %d\n", r.ExtraFindings)
	if len(r.HardMisses) > 0 {
		fmt.Fprintf(w, "\n**hard misses:**\n")
		for _, m := range r.HardMisses {
			fmt.Fprintf(w, "- %s\n", m)
		}
	}
	if len(r.SoftMisses) > 0 {
		fmt.Fprintf(w, "\n_soft misses:_\n")
		for _, m := range r.SoftMisses {
			fmt.Fprintf(w, "- %s\n", m)
		}
	}
}
