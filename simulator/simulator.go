package simulator

import (
	"math"

	"github.com/thoriqadillah/tal/entity"
)

type Simulator struct {
	Canvas *entity.Canvas
}

func (s *Simulator) Rule(p1 []entity.Particle, p2 []entity.Particle, gravity float64) {
	var fx, fy float64

	for i := range p1 {
		fx = 0
		fy = 0
		for j := range p2 {
			dx := p1[i].Position.X - p2[j].Position.X
			dy := p1[i].Position.Y - p2[j].Position.Y
			distance := math.Sqrt(float64(dx*dx + dy*dy))
			if distance > 0 && distance < 80 {
				F := gravity * 1 / distance
				fx += F * float64(dx)
				fy += F * float64(dy)
			}
		}
		p1[i].Vx = (p1[i].Vx + fx) * 0.6
		p1[i].Vy = (p1[i].Vy + fy) * 0.6

		s.Canvas.Cells[p1[i].Position.Y][p1[i].Position.X] = " "
		x := (p1[i].Position.X + int(math.Round(p1[i].Vx))) % s.Canvas.Width
		y := (p1[i].Position.Y + int(math.Round(p1[i].Vy))) % s.Canvas.Height

		if x <= 0 {
			x *= -1
		}

		if y <= 0 {
			y *= -1
		}

		p1[i].Position.X = x
		p1[i].Position.Y = y
		s.Canvas.Cells[p1[i].Position.Y][p1[i].Position.X] = p1[i].Char
	}
}
