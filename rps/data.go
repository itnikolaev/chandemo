package rps

import (
	"time"
)

const (
	flushInterval = 5
)

var rpsStore map[int64]int64
var rpsChannel chan int64
var timerChannel <-chan time.Time

func init() {
	rpsStore = make(map[int64]int64, flushInterval+1)
	rpsChannel = make(chan int64, 4096*3)
	go process()
}
