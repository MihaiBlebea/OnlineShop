package http

import (
	"math"
	"math/rand"
	"time"
)

func roundTwoDecimals(input float64) float64 {
	return math.Round(input*100) / 100
}

func genRandom(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
