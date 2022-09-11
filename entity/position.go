package entity

import "github.com/thoriqadillah/tal/helper"

type Position struct {
	X int
	Y int
}

func randomPosition(x int, y int) Position {
	return Position{
		X: helper.Random(x),
		Y: helper.Random(y),
	}
}
