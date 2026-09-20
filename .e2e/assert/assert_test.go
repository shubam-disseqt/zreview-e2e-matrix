// SPDX-License-Identifier: Apache-2.0
// Copyright 2026 shubam-ai-code-reviewer contributors

package main

import "testing"

func act(path, cat, sev string, start, end int) Actual {
	return Actual{Path: path, Category: cat, Severity: sev, StartLine: start, EndLine: end}
}
func exp(path, cat, sev string, start, end int) Expected {
	return Expected{Path: path, Category: cat, MinSeverity: sev, StartLine: start, EndLine: end}
}

func TestMatch_ExactLineExactCategory(t *testing.T) {
	if !Match(exp("a.go", "security", "high", 10, 12), act("a.go", "security", "high", 10, 12)) {
		t.Fatal("exact match should hit")
	}
}

func TestMatch_LineToleranceBothSides(t *testing.T) {
	if !Match(exp("a.go", "bug", "medium", 20, 25), act("a.go", "bug", "medium", 27, 27)) {
		t.Fatal("+2 line drift within tolerance should hit")
	}
	if Match(exp("a.go", "bug", "medium", 20, 25), act("a.go", "bug", "medium", 15, 16)) {
		t.Fatal("-4 line drift beyond tolerance should miss")
	}
}

func TestMatch_CategoryCaseInsensitive(t *testing.T) {
	if !Match(exp("a.go", "SECURITY", "high", 10, 10), act("a.go", "security", "high", 10, 10)) {
		t.Fatal("category compare should be case-insensitive")
	}
}

func TestMatch_WrongCategoryMisses(t *testing.T) {
	if Match(exp("a.go", "security", "high", 10, 10), act("a.go", "bug", "high", 10, 10)) {
		t.Fatal("wrong category should miss")
	}
}

func TestMatch_WrongPathMisses(t *testing.T) {
	if Match(exp("a.go", "bug", "low", 10, 10), act("b.go", "bug", "low", 10, 10)) {
		t.Fatal("wrong path should miss")
	}
}

func TestMatch_SeverityGate(t *testing.T) {
	if Match(exp("a.go", "bug", "medium", 10, 10), act("a.go", "bug", "low", 10, 10)) {
		t.Fatal("actual severity below floor should miss")
	}
	if !Match(exp("a.go", "bug", "medium", 10, 10), act("a.go", "bug", "high", 10, 10)) {
		t.Fatal("actual severity above floor should hit")
	}
}

func TestAssert_AllMatched(t *testing.T) {
	e := ExpectedFile{
		Case: "seeded-basic",
		Findings: []Expected{
			exp("a.go", "security", "high", 10, 10),
			exp("b.go", "bug", "medium", 20, 22),
		},
	}
	a := ActualFile{
		Comments: []Actual{
			act("a.go", "security", "high", 10, 10),
			act("b.go", "bug", "medium", 21, 21),
		},
	}
	r := Assert("seeded-basic", e, a)
	if !r.Passed || r.Matched != 2 || r.ExtraFindings != 0 || r.HardRecall != 1.0 {
		t.Errorf("expected clean pass, got %+v", r)
	}
}

func TestAssert_HardMissFails(t *testing.T) {
	e := ExpectedFile{
		Case:     "one-missed",
		Findings: []Expected{exp("a.go", "security", "high", 10, 10), exp("b.go", "bug", "medium", 20, 22)},
	}
	a := ActualFile{Comments: []Actual{act("a.go", "security", "high", 10, 10)}}
	r := Assert("one-missed", e, a)
	if r.Passed || len(r.HardMisses) != 1 || r.HardRecall != 0.5 {
		t.Errorf("expected fail with 1 hard miss, got %+v", r)
	}
}

func TestAssert_SoftMissDoesNotFail(t *testing.T) {
	e := ExpectedFile{
		Case: "soft-only",
		Findings: []Expected{
			exp("a.go", "security", "high", 10, 10),
			{Path: "b.go", Category: "maintainability", MinSeverity: "low", StartLine: 30, EndLine: 30, Soft: true},
		},
	}
	a := ActualFile{Comments: []Actual{act("a.go", "security", "high", 10, 10)}}
	r := Assert("soft-only", e, a)
	if !r.Passed || len(r.SoftMisses) != 1 || len(r.HardMisses) != 0 {
		t.Errorf("expected pass with 1 soft miss, got %+v", r)
	}
}

func TestAssert_MissedCriticalMarksSpecial(t *testing.T) {
	e := ExpectedFile{Case: "critical-missed", Findings: []Expected{exp("cfg.go", "security", "critical", 5, 5)}}
	r := Assert("critical-missed", e, ActualFile{})
	if !r.MissedCritical || r.Passed || r.CriticalSummary == "" {
		t.Fatalf("expected MissedCritical + fail + summary, got %+v", r)
	}
}

func TestAssert_ExtraFindingsCountedButNoFail(t *testing.T) {
	e := ExpectedFile{Case: "just-one", Findings: []Expected{exp("a.go", "bug", "medium", 10, 10)}}
	a := ActualFile{Comments: []Actual{act("a.go", "bug", "medium", 10, 10), act("z.go", "style", "low", 100, 100)}}
	r := Assert("just-one", e, a)
	if !r.Passed || r.ExtraFindings != 1 {
		t.Errorf("expected pass + 1 extra, got %+v", r)
	}
}

func TestAssert_EmptyExpectedPassesOnEmptyActual(t *testing.T) {
	r := Assert("docs-only", ExpectedFile{}, ActualFile{})
	if !r.Passed || r.HardRecall != 1.0 {
		t.Fatalf("empty in, empty out should pass with 100%% recall, got %+v", r)
	}
}

func TestAssert_EmptyExpectedWithExtrasStillPasses(t *testing.T) {
	r := Assert("docs-only", ExpectedFile{}, ActualFile{Comments: []Actual{act("README.md", "style", "low", 1, 1)}})
	if !r.Passed || r.ExtraFindings != 1 {
		t.Errorf("expected pass + 1 extra, got %+v", r)
	}
}
