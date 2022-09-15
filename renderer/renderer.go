package renderer

import (
	"time"

	"github.com/thoriqadillah/tal/simulator"
)

func Render(simulator *simulator.Simulator) {
	var cursor Cursor

	time.Sleep(1000 * time.Millisecond / 60) //this will make the renderer runs certain fps/s
	simulator.Canvas.Render()
	cursor.refresh()
}
