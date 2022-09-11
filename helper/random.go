package helper

import "math/rand"

func Random(max int) int {
	return rand.Intn(max-0) + 0
}
