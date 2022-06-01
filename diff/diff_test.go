package diff

import (
	"fmt"
	"testing"

	"github.com/api7/gotextdiff/difftest"
	"github.com/api7/gotextdiff/span"
)

func TestApplyEdits(t *testing.T) {
	for _, tc := range difftest.TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Helper()
			if got := ApplyEdits(tc.In, tc.Edits); got != tc.Out {
				t.Errorf("ApplyEdits edits got %q, want %q", got, tc.Out)
			}
			if tc.LineEdits != nil {
				if got := ApplyEdits(tc.In, tc.LineEdits); got != tc.Out {
					t.Errorf("ApplyEdits lineEdits got %q, want %q", got, tc.Out)
				}
			}
		})
	}
}

func TestLineEdits(t *testing.T) {
	for _, tc := range difftest.TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Helper()
			// if line edits not specified, it is the same as edits
			edits := tc.LineEdits
			if edits == nil {
				edits = tc.Edits
			}
			if got := LineEdits(tc.In, tc.Edits); diffEdits(got, edits) {
				t.Errorf("LineEdits got %q, want %q", got, edits)
			}
		})
	}
}

func TestUnified(t *testing.T) {
	for _, tc := range difftest.TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Helper()
			unified := fmt.Sprint(ToUnified(difftest.FileA, difftest.FileB, tc.In, tc.Edits))
			if unified != tc.Unified {
				t.Errorf("edits got diff:\n%v\nexpected:\n%v", unified, tc.Unified)
			}
			if tc.LineEdits != nil {
				unified := fmt.Sprint(ToUnified(difftest.FileA, difftest.FileB, tc.In, tc.LineEdits))
				if unified != tc.Unified {
					t.Errorf("lineEdits got diff:\n%v\nexpected:\n%v", unified, tc.Unified)
				}
			}
		})
	}
}

func diffEdits(got, want []TextEdit) bool {
	if len(got) != len(want) {
		return true
	}
	for i, w := range want {
		g := got[i]
		if span.Compare(w.Span, g.Span) != 0 {
			return true
		}
		if w.NewText != g.NewText {
			return true
		}
	}
	return false
}
