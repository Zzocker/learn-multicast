package algo

import (
	"log"
	"os"

	"github.com/Zzocker/multicast/algo/antientropy"
	"github.com/Zzocker/multicast/algo/msg"
	"github.com/Zzocker/multicast/utils"
)

type ALGO_NAME int

const (
	ALGO_ANTI_ENTROPY ALGO_NAME = iota + 1
	ALGO_RUMOR_MONGERING
	ALGO_GOSSIP
)

type Peer interface {
	Start(neighbor []chan msg.MulticastMessage)
	Set(value int64)
	Get() int64
}

type Simulator struct {
	lg    *log.Logger
	peers []Peer
}

func (s *Simulator) Start(algoName ALGO_NAME, peerCount int) {
	s.lg = log.New(os.Stdout, "[SIMULATOR] ", 0)
	s.lg.Printf("algoName=%d, peerCount=%d", algoName, peerCount)

	chans := make([]chan msg.MulticastMessage, peerCount)
	chanBufferSize := 64

	for i, _ := range chans {
		chans[i] = make(chan msg.MulticastMessage, chanBufferSize)
	}

	switch algoName {
	case ALGO_ANTI_ENTROPY:
		for i := 0; i < peerCount; i++ {
			s.peers = append(s.peers, antientropy.New(i, chans[i]))
		}
	}

	// overlay setup
	aux := make([]bool, peerCount)
	for i := range s.peers {
		neighborI := neighborSelection(i, aux)
		var neighbor []chan msg.MulticastMessage
		for _, i := range neighborI {
			neighbor = append(neighbor, chans[i])
		}
		s.lg.Printf("id=%d ==> neighbors %v", i, neighborI)
		go s.peers[i].Start(neighbor)
	}

	for {
	}
}

// neighborSelection returns array of index of its neighboring peers
func neighborSelection(id int, aux []bool) []int {
	peerCount := len(aux)
	count := utils.RandBetween(1, (peerCount/2)+1)
	aux[id] = true
	var out []int
	for count > 0 {
		n := utils.RandBetween(0, peerCount)
		if !aux[n] {
			out = append(out, n)
			aux[n] = true
			count--
		}
	}
	for i := range aux {
		aux[i] = false
	}
	return out
}

func (s *Simulator) Set(value int64) {
	id := utils.RandBetween(0, len(s.peers))
	s.lg.Printf("Set call to peer %d", id)
	s.peers[id].Set(value)
}

func (s *Simulator) Get() int64 {
	id := utils.RandBetween(0, len(s.peers))
	s.lg.Printf("Get call to peer %d", id)
	return s.peers[id].Get()
}
