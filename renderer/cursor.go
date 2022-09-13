package renderer

import "fmt"

type Cursor struct{}

func (c *Cursor) moveToBeginning() {
	fmt.Print("\u001b[H")
}
