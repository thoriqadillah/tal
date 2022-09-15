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
	green := canvas.DrawParticle(20, color.Green("o"))
	red := canvas.DrawParticle(100, color.Red("o"))
	blue := canvas.DrawParticle(50, color.Blue("o"))

	simulator := simulator.Simulator{
		Canvas: canvas,
	}

	for {
		// simulator.Rule(red, green, -0.15)
		simulator.Rule(red, red, 0.05, 1)
		simulator.Rule(red, blue, -1, 0.6)

		simulator.Rule(green, green, 0.1, 1)
		simulator.Rule(green, red, -0.31, 0.6)
		simulator.Rule(green, blue, 0.14, 1)

		simulator.Rule(blue, blue, 0.15, 1)
		simulator.Rule(blue, green, -0.15, 0.9)
		simulator.Rule(blue, red, -0.25, 0.8)

		renderer.Render(&simulator)
	}
}
