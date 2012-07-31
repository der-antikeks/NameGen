// Package rand implements a weighted random selection.
package rand

import (
	"time"
	"math/rand"
)

// Seed seeds math.Seed(seed int64)) with time.Now().
func Seed() {
	rand.Seed(time.Now().UnixNano())
}

// Sum calculates the sum of a slice.
func Sum(weights []float64) float64 {
	sum := 0.0
	for _, w := range(weights) {
		sum += w
	}

	return sum
}

// WeightedChoice selects an random key from a slice.
// The probability of each key is based on the value.
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
