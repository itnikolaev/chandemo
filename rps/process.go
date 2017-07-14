package rps

import (
	"time"
)

func process() {

	var when int64
	db := NewFakeDatabase()
	timerChannel = time.Tick(time.Second * flushInterval)
	for {
		select {
		case when = <-rpsChannel:
			//log.Printf("rps, when: %d\n", when)
			rpsStore[when]++
		case <-timerChannel:
			go db.Store(rpsStore)
			rpsStore = make(map[int64]int64, flushInterval+1)
		}
	}
}
