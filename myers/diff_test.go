// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package myers_test

import (
	"testing"

	"github.com/api7/gotextdiff/diff"
	"github.com/api7/gotextdiff/myers"
)

func TestDiff(t *testing.T) {
	diff.DiffTest(t, myers.ComputeEdits)
}
