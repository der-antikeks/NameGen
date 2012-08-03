/*
	Package ssyll implements a Simple Syllable Divider.
	Usually works but not always correctly.
*/
package ssyll

import (
	"strings"
	"regexp"
)

// Simple english vowels, ignore Y as only-sometimes-vowel
const Vowels = "aeiou"

// Divide splits a word into its vowels
func Divide(word string) (syllables []string) {
	word = strings.ToLower(word)
	var lv, rv, dist int
	
	for len(word) > 0 {
		lv = strings.IndexAny(word, Vowels)
		if lv == -1 {
			syllables = append(syllables, word)
			break
		}

		rv = strings.IndexAny(word[lv+1:], Vowels)
		for rv == 0 {
			lv++
			rv = strings.IndexAny(word[lv+1:], Vowels)
		}

		if rv == -1 {
			syllables = append(syllables, word)
			break
		}

		dist = ((rv + lv) - lv) / 2
		if dist > -1 {
			syllables = append(syllables, word[:lv+1+dist])
			word = word[lv+1+dist:]
		}
	}

	return
}

// DivideText splits a full text into separate words and those into its vowels
func DivideText(text string) [][]string {
	ret := [][]string{}
	text = regexp.MustCompile("[^a-zA-Z]+").ReplaceAllString(text, " ")

	words := strings.Split(text, " ")
	for _, word := range words {
		ret = append(ret, Divide(word))
	}

	return ret
}
