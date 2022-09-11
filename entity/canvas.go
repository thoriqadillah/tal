package entity

import "fmt"

type Canvas struct {
	Width  int
	Height int
	Cells  Cell
}

func (c *Canvas) DrawCanvas() *Canvas {
	cells := make(Cell, c.Height)
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
	particles := make([]Particle, total)

	for i := 0; i < total; i++ {
		particles[i] = Particle{
			Position: randomPosition(c.Width, c.Height),
			Char:     char,
		}

		x := particles[i].Position.X
		y := particles[i].Position.Y

		c.Cells[y][x] = particles[i].Char
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
