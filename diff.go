package gotextdiff

import (
	"github.com/api7/gotextdiff/diff"
	"github.com/api7/gotextdiff/myers"
	"github.com/api7/gotextdiff/span"
)

func StringDiff(str1, str2 string, options ...diff.Option) diff.Unified {
	edits := myers.ComputeEdits(span.URIFromPath("t"), str1, str2)

	return diff.ToUnified("", "", str1, edits, options...)
}
