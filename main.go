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
	green := canvas.DrawParticle(20, color.Green("O"))
	red := canvas.DrawParticle(30, color.Red("O"))
	blue := canvas.DrawParticle(50, color.Blue("O"))

	simulator := simulator.Simulator{
		Canvas: canvas,
	}

	for {
		// simulator.Rule(red, green, -0.15)
		simulator.Rule(red, red, 0.05)
		simulator.Rule(red, blue, -1)

		simulator.Rule(green, green, 0.1)
		simulator.Rule(green, red, -0.31)
		simulator.Rule(green, blue, 0.14)

		simulator.Rule(blue, blue, 0.15)
		simulator.Rule(blue, green, -0.15)
		simulator.Rule(blue, red, -0.25)

		renderer.Render(&simulator)
	}
}
