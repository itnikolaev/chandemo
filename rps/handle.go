package rps

func AddMetric(when int64) {
	rpsChannel <- when
}
