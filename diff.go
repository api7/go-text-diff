package gotextdiff

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/api7/gotextdiff/diff"
	"github.com/api7/gotextdiff/myers"
	"github.com/api7/gotextdiff/span"
)

func StringsDiff(str1, str2 string, options ...diff.Option) *diff.Unified {
	edits := myers.ComputeEdits(span.URIFromPath(""), str1, str2)
	unified := diff.ToUnified("", "", str1, edits, options...)

	return &unified
}

func FilesDiff(fileFrom, fileTo string, options ...diff.Option) *diff.Unified {
	content1, err := ioutil.ReadFile(fileFrom)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file %s: %s", fileFrom, err)
		os.Exit(1)
		return nil
	}
	content2, err := ioutil.ReadFile(fileTo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file %s: %s", fileTo, err)
		os.Exit(1)
		return nil
	}

	edits := myers.ComputeEdits(span.URIFromPath(fileFrom), string(content1), string(content2))
	unified := diff.ToUnified(fileFrom, fileTo, string(content1), edits, options...)
	return &unified
}
