package utils

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkRandBetween(b *testing.B) {
	rand.Seed(time.Now().Unix())
	low := 34
	high := 65
	for i := 0; i < b.N; i++ {
		v := RandBetween(low, high)
		if !(v >= low && v < high) {
			b.Fatalf("value %d is not in range [%d,%d)", v, low, high)
		}
	}
}
