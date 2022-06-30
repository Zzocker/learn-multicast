package antientropy

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Zzocker/multicast/algo/msg"
	"github.com/Zzocker/multicast/utils"
)

type AntiEntropy struct {
	lg         *log.Logger
	value      int64
	timestamp  int64
	mux        *sync.Mutex
	listenChan <-chan msg.MulticastMessage
}

func New(id int, listenChan <-chan msg.MulticastMessage) *AntiEntropy {
	return &AntiEntropy{
		lg:         log.New(os.Stdout, fmt.Sprintf("[peer-%d]", id), 0),
		value:      -1,
		timestamp:  0,
		listenChan: listenChan,
		mux:        &sync.Mutex{},
	}
}

func (a *AntiEntropy) Start(neighbor []chan msg.MulticastMessage) {
	a.lg.Println("start")
	go a.listener()
	for {
		time.Sleep(time.Millisecond * time.Duration(utils.RandBetween(400, 600)))
		id := utils.RandBetween(0, len(neighbor))
		a.mux.Lock()
		msg := msg.MulticastMessage{
			Data:      a.value,
			Timestamp: a.timestamp,
		}
		a.mux.Unlock()
		neighbor[id] <- msg
	}
}

func (a *AntiEntropy) listener() {
	for msg := range a.listenChan {
		a.mux.Lock()
		// push based anti-entropy protocol
		if msg.Timestamp > a.timestamp {
			a.value = msg.Data
			a.timestamp = msg.Timestamp
			a.lg.Printf("State: value=%d, timestamp=%d", msg.Data, msg.Timestamp)
		}
		a.mux.Unlock()
	}
}

func (a *AntiEntropy) Set(value int64) {
	a.lg.Printf("Set %d", value)
	a.mux.Lock()
	defer a.mux.Unlock()
	a.timestamp = time.Now().Unix()
	a.value = value
}

func (a *AntiEntropy) Get() int64 {
	a.lg.Println("Get")
	a.mux.Lock()
	defer a.mux.Unlock()
	return a.value
}
