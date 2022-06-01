# gotextdiff - unified text diffing in Go

This is a copy of the Go text diffing packages that [the official Go language server gopls uses internally](https://github.com/golang/tools/tree/master/internal/lsp/diff) to generate unified diffs.

If you've previously tried to generate unified text diffs in Go (like the ones you see in Git and on GitHub), you may have found [github.com/sergi/go-diff](https://github.com/sergi/go-diff) which is a Go port of Neil Fraser's google-diff-match-patch code - however it [does not support unified diffs](https://github.com/sergi/go-diff/issues/57).

This is arguably one of the best (and most maintained) unified text diffing packages in Go as of at least 2020.

(All credit goes to [the Go authors](http://tip.golang.org/AUTHORS), We are merely re-publishing their work so others can use it.)

## Example usage

Assuming you want to diff `a.txt` and `b.txt`, whose contents are stored in `aString` and `bString` then:

```Go
import (
    "fmt"

    "github.com/api7/gotextdiff"
    "github.com/api7/gotextdiff/myers"
)

aString := `foo`
bString := `bar`
edits := myers.ComputeEdits(span.URIFromPath("a.txt"), aString, bString)

fmt.Println(gotextdiff.ToUnified("a.txt", "b.txt", aString, edits))
```

the output will be a string like:

```diff
--- a.txt
+++ b.txt
@@ -1,13 +1,28 @@
-foo
+bar
```
