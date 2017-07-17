package strategy

import (
	"fmt"

	"github.com/Zenithar/typogenerator/helpers"
)

var (
	BitSquatting Strategy
)

type bitsquattingStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *bitsquattingStrategy) Generate(dom string) ([]string, error) {
	res := []string{}

	masks := []int32{1, 2, 4, 8, 16, 32, 64, 128}

	for i, c := range dom {
		for m := range masks {
			b := rune(int(c) ^ m)
			o := int(b)
			if (o >= 48 && o <= 57) || (o >= 97 && o <= 122) || o == 45 {
				res = append(res, fmt.Sprintf("%s%c%s", dom[:i], b, dom[i+1:]))
			}
		}
	}

	return helpers.Dedup(res), nil
}

func (s *bitsquattingStrategy) GetName() string {
	return "BitSquatting"
}

func init() {
	BitSquatting = &bitsquattingStrategy{}
}
