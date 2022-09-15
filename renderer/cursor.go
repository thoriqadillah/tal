package renderer

import "fmt"

type Cursor struct{}

func (c *Cursor) refresh() {
	fmt.Print("\u001b[H")
}
