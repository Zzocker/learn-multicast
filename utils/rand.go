package utils

import (
	"math/rand"
)

func RandBetween(low, high int) int {
	return rand.Intn(high-low) + low
}
