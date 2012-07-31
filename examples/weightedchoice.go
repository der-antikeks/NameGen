package main

import (
	"fmt"

	"github.com/der-antikeks/namegen/rand"
)

func main() {
	data 	:= []string {"a", "b", "c"}
	weights	:= []float64{0.6, 0.3, 0.1}
	results := map[string]int {"a": 0, "b": 0, "c": 0}
	runs	:= 1000

	fmt.Printf("%v => %v\n", data, weights)

	rand.Seed()
	for i := 0; i < runs; i++ {
		index := rand.WeightedChoice(weights, rand.Sum(weights))
		results[data[index]]++
	}

	for k, v := range(results) {
		fmt.Printf("%v => %v\n", k, float64(v)/float64(runs))
	}
}
