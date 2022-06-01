package main

import (
	"fmt"

	"github.com/api7/gotextdiff"
	"github.com/api7/gotextdiff/diff"
)

func main() {
	fmt.Print(gotextdiff.FilesDiff("./example/fixtures/config.yaml", "./example/fixtures/config2.yaml",
		diff.Colorful(true), diff.OmitEOL(true)))
}
