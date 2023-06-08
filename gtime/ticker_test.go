package gtime

import (
	"testing"
	"time"
)

func TestRealTickerDoTick(t *testing.T) {
	ticker := NewRealTicker(time.Millisecond * 10)
	defer ticker.Closed()
	var count int
	for range ticker.Chan() {
		count++
		if count > 5 {
			break
		}
	}
}
