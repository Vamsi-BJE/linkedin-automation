package stealth

import (
	"math/rand"
	"time"
)

func Delay(minMs, maxMs int) {
	if maxMs <= minMs {
		maxMs = minMs + 100
	}
	time.Sleep(time.Duration(rand.Intn(maxMs-minMs)+minMs) * time.Millisecond)
}
