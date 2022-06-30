package algo

import (
	"math/rand"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	s := &Simulator{}

	go s.Start(ALGO_ANTI_ENTROPY, 8)

	time.Sleep(time.Second)
	s.Set(10)
	time.Sleep(time.Second * 5)
	t.Log(s.Get())
}

func TestNeighborSelection(t *testing.T) {
	rand.Seed(time.Now().Unix())
	peerCount := 8
	aux := make([]bool, peerCount)

	neg := neighborSelection(2, aux)
	t.Log(neg)
	t.Log(aux)
}
