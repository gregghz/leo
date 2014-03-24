package golang

import (
	"../node"
	"io"
	"fmt"
)

func Gen(root node.Node, out io.Writer) {
	fmt.Fprintf(out, root.Go())
}
