package main

import (
	"github.com/thoriqadillah/tal/entity"
	"github.com/thoriqadillah/tal/renderer"
	"github.com/thoriqadillah/tal/simulator"
)

func main() {
	var color renderer.Color

	c := entity.Canvas{
		// Width:  150,
		// Height: 40,
		Width:  270,
		Height: 70,
	}

	canvas := c.DrawCanvas()
	round := canvas.DrawParticle(2, color.Green("O"))

	simulator := simulator.Simulator{
		Canvas: canvas,
	}

	for {
		simulator.Rule(round, round, -1)
		renderer.Render(&simulator)
	}
}
