// SPDX-License-Identifier: Apache-2.0
// Copyright 2026 shubam-ai-code-reviewer contributors

package main

import "testing"

func mkResult(name string, passed bool, expected, matched, extras int) Result {
	r := Result{
		Case:          name,
		ExpectedHard:  expected,
		MatchedHard:   matched,
		ExpectedTotal: expected,
		Matched:       matched,
		ExtraFindings: extras,
		Passed:        passed,
	}
	if expected > 0 {
		r.HardRecall = float64(matched) / float64(expected)
	} else {
		r.HardRecall = 1.0
	}
	return r
}

func TestAggregate_AllPass(t *testing.T) {
	s := Aggregate([]Result{
		mkResult("secret", true, 1, 1, 0),
		mkResult("sql", true, 2, 2, 1),
	}, 0.8)
	if !s.Passed || s.HardRecall != 1.0 || s.ExtraFindingsTotal != 1 {
		t.Errorf("expected clean 100%% pass, got %+v", s)
	}
}

func TestAggregate_BelowThresholdFails(t *testing.T) {
	s := Aggregate([]Result{mkResult("a", true, 2, 2, 0), mkResult("b", false, 2, 0, 0)}, 0.8)
	if s.Passed || s.HardRecall != 0.5 {
		t.Fatalf("expected fail at 50%%, got %+v", s)
	}
}

func TestAggregate_AtThresholdPasses(t *testing.T) {
	s := Aggregate([]Result{mkResult("a", true, 2, 2, 0), mkResult("b", true, 3, 2, 0)}, 0.8)
	if !s.Passed {
		t.Fatalf("expected pass at 80%%, got %+v", s)
	}
}

func TestAggregate_MissedCriticalFailsEvenAboveThreshold(t *testing.T) {
	r := mkResult("secret", false, 1, 0, 0)
	r.MissedCritical = true
	r.CriticalSummary = "config.go:5-5 [security/critical]"
	s := Aggregate([]Result{r, mkResult("a", true, 20, 20, 0)}, 0.8)
	if s.Passed || len(s.MissedCriticalIn) != 1 {
		t.Fatalf("missed critical should fail the whole matrix, got %+v", s)
	}
}

func TestAggregate_EmptyExpectedResultsPass(t *testing.T) {
	s := Aggregate([]Result{mkResult("docs", true, 0, 0, 2), mkResult("real", true, 3, 3, 0)}, 0.8)
	if !s.Passed || s.HardRecall != 1.0 || s.ExtraFindingsTotal != 2 {
		t.Errorf("expected pass with 100%% and 2 extras, got %+v", s)
	}
}

func TestAggregate_ZeroHardExpectedYields100(t *testing.T) {
	s := Aggregate([]Result{mkResult("docs", true, 0, 0, 0)}, 0.8)
	if !s.Passed || s.HardRecall != 1.0 {
		t.Fatalf("zero-expected matrix should pass at 100%%, got %+v", s)
	}
}
