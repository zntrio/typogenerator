package strategy

import (
	"fmt"

	"go.zenithar.org/typogenerator/helpers"
)

var (
	BitSquatting Strategy
)

type bitsquattingStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *bitsquattingStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	masks := []int32{1, 2, 4, 8, 16, 32, 64, 128}

	for i, c := range domain {
		for m := range masks {
			b := rune(int(c) ^ m)
			o := int(b)
			if (o >= 48 && o <= 57) || (o >= 97 && o <= 122) || o == 45 {
				fuzzed := fmt.Sprintf("%s%c%s", domain[:i], b, domain[i+1:])
				fuzzed = combineTLD(fuzzed, tld)
				res = append(res, fuzzed)
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
