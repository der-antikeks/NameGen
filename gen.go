// Package namegen creates random names based on the probability of characters in a given string.
package namegen

import (
	"strings"

	"github.com/der-antikeks/namegen/rand"
	"github.com/der-antikeks/namegen/ssyll"
)


const (
	StopString = " "	// The string that is used to separate names.
	Vowels = "aeiou"	// Simple english vowels for syllable splitting, ignore Y as only-sometimes-vowel
)

// A NameGen creates random names.
type NameGen struct {
	dict	map[string]map[string]int	// [previous syllable][next syllable]amount
	length	map[int]int					// [number of syllables]amount
}

// NewNameGen returns a new NameGen and calculates the probabilities of consecutive syllables based on the passed string slice.
func NewNameGen(names []string) *NameGen {
	n := new(NameGen)

	n.length = make(map[int]int)
	n.dict = make(map[string]map[string]int)

	var syllables []string

	prev := StopString
	for _, name := range names {
		name = strings.ToLower(name)
		syllables = ssyll.Divide(name, Vowels)

		n.length[len(syllables)]++

		for _, syllable := range syllables {
			if n.dict[prev] == nil {
				n.dict[prev] = make(map[string]int)
			}

			n.dict[prev][syllable]++
			prev = syllable
		}

		if n.dict[prev] == nil {
			n.dict[prev] = make(map[string]int)
		}
		n.dict[prev][StopString]++

		prev = StopString
	}

	return n
}

// GenerateOne generates a single name.
func (n NameGen) GenerateOne() string {
	return n.GenerateWithStart(string(StopString))
}

// GenerateMultiple generates multiple names.
func (n NameGen) GenerateMultiple(amount int) []string {
	ret := make([]string, amount)
	for i := range ret {
		ret[i] = n.GenerateOne()
	}
	
	return ret
}

// GenerateWithStart generates a single name with a specified start string.
func (n NameGen) GenerateWithStart(start string) string {
	var cur string
	var name []string
	
	prev := start
	
	for (cur != StopString) {
		if n.dict[prev] == nil {
			break
		}
		
		cur = selectString(n.dict[prev])
		name = append(name, cur)
		prev = cur
	}
	
	return strings.TrimSpace(strings.Join(name, ""))
}

// selectString chooses a string based on the weight.
func selectString(dict map[string]int) string {
	rand.Seed()

	data 	:= []string{}
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
