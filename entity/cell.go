package entity

import "github.com/thoriqadillah/tal/helper"

type Cell struct {
	X int
	Y int
}

func randomPosition(x int, y int) Cell {
	return Cell{
		X: helper.Random(x),
		Y: helper.Random(y),
	}
}
