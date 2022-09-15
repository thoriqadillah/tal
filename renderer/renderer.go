package renderer

import (
	"time"

	"github.com/thoriqadillah/tal/simulator"
)

func Render(simulator *simulator.Simulator) {
	var cursor Cursor

	time.Sleep(1000 * time.Millisecond / 30)
	simulator.Canvas.Render()
	cursor.refresh()
}
