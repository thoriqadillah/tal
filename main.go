package main

import (
	"github.com/thoriqadillah/tal/entity"
)

func main() {
	c := entity.Canvas{
		// Width:  150,
		// Height: 40,
		Width:  270,
		Height: 70,
	}

	canvas := c.DrawCanvas()
	canvas.DrawParticle(10, ".").DrawParticle(10, "0").Render()
}
