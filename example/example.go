package main

import (
	"fmt"

	"github.com/api7/gotextdiff"
	"github.com/api7/gotextdiff/diff"
)

func main() {

	aString := `The Way that can be told of is not the eternal Way;
The name that can be named is not the eternal name.
The Nameless is the origin of Heaven and Earth;
The Named is the mother of all things.
Therefore let there always be non-being,
  so we may see their subtlety,
And let there always be being,
  so we may see their outcome.
The two are the same,
But after they are produced,
  they have different names.`

	bString := `The Way that can be told of is not the eternal Way;
The name that can be named is not the eternal name.
The Nameless is th11e origin of Heaven and Earth;
The Named is the mother of all things.
Therefore let there always be non-being,
  so we may see their subtlety,
And let there always be being,
  so we may see their outcome.
The two are the same,
But after they are produced,
  they have different22 names.`

	fmt.Print(gotextdiff.StringsDiff(aString, bString,
		diff.Colorful(true), diff.OmitEOL(true)))

	fmt.Println("")

	fmt.Print(gotextdiff.FilesDiff("./example/fixtures/config.yaml", "./example/fixtures/config2.yaml",
		diff.Colorful(true), diff.OmitEOL(true)))

}
