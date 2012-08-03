package main

import (
	"fmt"
	"strings"

	"github.com/der-antikeks/namegen/ssyll"
)

func main() {
	words := ssyll.DivideText("bAcdEfghI, AEhjklO, vwEgkldcIkd. jAcOb, mAsOn, wIllIAm, jAydEn, nOAh, mIchAEl, EthAn, AlExAndEr, IsAAc, gAvIn, brAydEn, tYlEr. ErgOnOmIc, mOUntAInOUs, vIllAInOUs, bAchElOr, gEnUInE, mAstOdOn, sqUIrmEd.")
	for _, syllables := range words {
		fmt.Printf("%v => %v\n", strings.Join(syllables, ""), strings.Join(syllables, "|"))
	}
}
