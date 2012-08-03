package main

import (
	"fmt"
	"strings"

	"github.com/der-antikeks/namegen/ssyll"
)

func main() {
	vowels := "aeiou"
	words := ssyll.DivideText("bAcdEfghI, AEhjklO, vwEgkldcIkd. jAcOb, mAsOn, wIllIAm, jAydEn, nOAh, mIchAEl, EthAn, AlExAndEr, IsAAc, gAvIn, brAydEn, tYlEr. ErgOnOmIc, mOUntAInOUs, vIllAInOUs, bAchElOr, gEnUInE, mAstOdOn, sqUIrmEd.", vowels)
	for _, syllables := range words {
		fmt.Printf("%v => %v\n", strings.Join(syllables, ""), strings.Join(syllables, "|"))
	}
}
