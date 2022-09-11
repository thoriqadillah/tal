package entity

import "fmt"

type Canvas struct {
	Width  int
	Height int
	Cells  [][]string
}

func (c *Canvas) DrawCanvas() *Canvas {
	cells := make([][]string, c.Height)
	for i := range cells {
		cells[i] = make([]string, c.Width)
	}

	for i := 0; i < c.Height; i++ {
		for j := 0; j < c.Width; j++ {
			cells[i][j] = " "
		}
	}

	c.Cells = cells
	return c
}

func (c *Canvas) DrawParticle(total int, char string) *Canvas {
	particleGroup := make([]Particle, total)

	for i := 0; i < total; i++ {
		particleGroup[i] = Particle{
			Position: randomPosition(c.Width, c.Height),
			Char:     char,
		}

		x := particleGroup[i].Position.X
		y := particleGroup[i].Position.Y

		c.Cells[y][x] = particleGroup[i].Char
	}

	return c
}

func (c *Canvas) Render() {
	for i := 0; i < c.Height; i++ {
		for j := 0; j < c.Width; j++ {
			fmt.Printf("%+v", c.Cells[i][j])
		}
		println()
	}
}
