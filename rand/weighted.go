package rand

import (
	"time"
	"math/rand"
)

func Seed() {
	rand.Seed(time.Now().UnixNano())
}

func Sum(weights []float64) float64 {
	sum := 0.0
	for _, w := range(weights) {
		sum += w
	}

	return sum
}

func WeightedChoice(weights []float64, sum float64) int {
	rnd := rand.Float64() * sum
	for i, w := range(weights) {
		rnd -= w

		if rnd < 0 {
			return i
		}
	}
	
	return 0
}
