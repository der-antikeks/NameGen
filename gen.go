// Package namegen creates random names based on the probability of characters in a given string.
package namegen

import (
	"strings"

	"github.com/der-antikeks/namegen/rand"
)

// The rune that is used to separate names.
const StopChar = rune(32)

// A NameGen creates random names.
type NameGen struct {
	dict	map[rune]map[rune]int	// [previous char][next char]amount
	length	map[int]int				// [length]amount
}

// NewNameGen returns a new NameGen and calculates the probabilities of consecutive characters based on the passed string slice.
func NewNameGen(raw []string) *NameGen {
	n := new(NameGen)

	n.length = make(map[int]int)
	n.dict = make(map[rune]map[rune]int)

	prev := StopChar
	for _, name := range raw {
		name = strings.ToLower(name)
		n.length[len(name)]++

		for _, char := range name {
			if n.dict[prev] == nil {
				n.dict[prev] = make(map[rune]int)
			}

			n.dict[prev][char]++
			prev = char
		}

		if n.dict[prev] == nil {
			n.dict[prev] = make(map[rune]int)
		}
		n.dict[prev][StopChar]++

		prev = StopChar
	}

	return n
}

// GenerateOne generates a single name.
func (n NameGen) GenerateOne() string {
	return n.GenerateWithStart(string(StopChar))
}

// GenerateMultiple generates multiple names.
func (n NameGen) GenerateMultiple(amount int) []string {
	ret := make([]string, amount)
	for i := range ret {
		ret[i] = n.GenerateOne()
	}
	
	return ret
}

// GenerateWithStart generates a single name with a specified start character.
func (n NameGen) GenerateWithStart(start string) string {
	var cur rune
	var name []rune
	
	prev := rune(start[0])
	
	for (cur != StopChar) {
		if n.dict[prev] == nil {
			break
		}
		
		cur = selectRune(n.dict[prev])
		name = append(name, cur)
		prev = cur
	}
	
	return strings.TrimSpace(string(name))
}

// selectRune chooses a rune based on the weight.
func selectRune(dict map[rune]int) rune {
	rand.Seed()

	data 	:= []rune{}
	weights	:= []float64{}
	sum 	:= 0.0

	for _, w := range dict {
		sum += float64(w)
	}

	for c, w := range dict {
		data = append(data, c)
		weights = append(weights, float64(w) / sum)
	}
	index := rand.WeightedChoice(weights, sum)
	result := data[index]

	return result
}
