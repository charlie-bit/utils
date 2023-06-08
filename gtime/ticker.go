package gtime

import "time"

type (
	Ticker interface {
		Chan() <-chan time.Time
		Closed()
	}

	realTicker struct {
		*time.Ticker
	}
)

func NewRealTicker(d time.Duration) Ticker {
	return &realTicker{
		Ticker: time.NewTicker(d),
	}
}

func (rt *realTicker) Chan() <-chan time.Time {
	return rt.C
}

func (rt *realTicker) Closed() {
	rt.Stop()
}
