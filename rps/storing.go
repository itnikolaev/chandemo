package rps

type RPSDatabase interface {
	Store(map[int64]int64) error
}
