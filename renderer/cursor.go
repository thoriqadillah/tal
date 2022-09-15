package renderer

import "fmt"

type Cursor struct{}

// set cursor of the terminal to top left corner to refresh and update the canvas
func (c *Cursor) refresh() {
	fmt.Print("\u001b[H")
}
