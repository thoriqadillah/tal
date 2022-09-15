package simulator

import (
	"math"

	"github.com/thoriqadillah/tal/entity"
)

type Simulator struct {
	Canvas *entity.Canvas
}

func (s *Simulator) Rule(p1 []entity.Particle, p2 []entity.Particle, gravity float64, acceleration float64) {
	var fx, fy float64

	dx := p1[0].Position.X - p2[1].Position.X
	dy := p1[0].Position.Y - p2[1].Position.Y
	distance := math.Sqrt(float64(dx*dx + dy*dy))
	if distance > 0 {
		F := gravity * 1 / distance
		fx += F * float64(dx)
		fy += F * float64(dy)
	}

	s.Canvas.Cells[p1[0].Position.Y][p1[0].Position.X] = " "
	p1[0].Position.X += int(math.Ceil(fx))
	p1[0].Position.Y += int(math.Ceil(fy))
	s.Canvas.Cells[p1[0].Position.Y][p1[0].Position.X] = p1[0].Char
}
