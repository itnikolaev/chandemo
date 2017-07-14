package rps

import (
	"log"
	"runtime"
)

type FakeDatabase struct {
}

func NewFakeDatabase() RPSDatabase {
	return new(FakeDatabase)
}

func (this *FakeDatabase) Store(m map[int64]int64) error {
	log.Printf("%+v", m)
	log.Printf("process: %d\n", runtime.GOMAXPROCS(-1))
	return nil
}
