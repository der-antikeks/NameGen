package namegen

import (
	"fmt"
	"strings"
)

const StopChar = rune(32)

type NameGen struct {
	dict	map[rune]map[float64]rune	// [previous char][probability]next char
	length	map[float64]int				// [probability]length
}

func NewNameGen(raw []string) *NameGen {
	n := new(NameGen)

	tmplength := make(map[int]int)
	tmpdict := make(map[rune]map[rune]int)
	prev := StopChar

	for _, name := range raw {
		name = strings.ToLower(name)
		tmplength[len(name)]++

		for _, char := range name {
			if tmpdict[prev] == nil {
				tmpdict[prev] = make(map[rune]int)
			}

			tmpdict[prev][char]++
			prev = char
		}

		if tmpdict[prev] == nil {
			tmpdict[prev] = make(map[rune]int)
		}
		tmpdict[prev][StopChar]++

		prev = StopChar
	}

	// char to char probability
	sum := 0
	sumchar := make(map[rune]float64)

	for prevchar, nextmap := range tmpdict {
		sum = 0
		for _, cnt := range nextmap {
			sum += cnt
		}

		sumchar[prevchar] = float64(sum)
	}

	n.dict = make(map[rune]map[float64]rune)
	for prevchar, nextmap := range tmpdict {
		sum = 0
		for nextchar, cnt := range nextmap {
			prob := float64(cnt) / sumchar[prevchar]

			if n.dict[prevchar] == nil {
				n.dict[prevchar] = make(map[float64]rune)
			}

			n.dict[prevchar][prob] = nextchar
		}
	}
	// TODO: pre-sort nextchar-probability desc

	// name length probability
	sum = 0
	for _, cnt := range tmplength {
		sum += cnt
	}

	n.length = make(map[float64]int, len(tmplength))
	for lng, cnt := range tmplength {
		prob := float64(cnt) / float64(sum)
		n.length[prob] = lng
	}
	// TODO: pre-sort length-probability desc

	return n
}

func (n NameGen) GenerateOne() string {
	return n.GenerateWithStart(string(StopChar))
}

func (n NameGen) GenerateMultiple(amount int) []string {
	ret := make([]string, amount)
	for i := range(ret) {
		ret[i] = n.GenerateOne()
	}
	
	return ret
}

func (n NameGen) GenerateWithStart(start string) string {
	var cur rune
	var name []rune
	
	prev := rune(start[0])
	
	for (len(name) <= 0 && cur != StopChar) {
		fmt.Printf("prev: [%c]\n", prev)
		
		if n.dict[prev] == nil {
			fmt.Printf("char %v not found\n", prev)
			break
		}
		
		cur = n.dict[prev][/* TODO: insert weighted-random function here */]
		
		name = append(name, cur)
		prev = cur
	}
	
	return strings.TrimSpace(string(name))
}
