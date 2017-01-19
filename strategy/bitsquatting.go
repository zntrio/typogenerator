package strategy

import (
	"fmt"
	"strings"

	"zenithar.org/go/typogenerator/helpers"
)

var (
	BitSquatting Strategy
)

type bitsquattingStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *bitsquattingStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := parts[0]
	tld := parts[1]
	masks := []int32{1, 2, 4, 8, 16, 32, 64, 128}

	for i, c := range dom {
		for m := range masks {
			b := rune(int(c) ^ m)
			o := int(b)
			if (o >= 48 && o <= 57) || (o >= 97 && o <= 122) || o == 45 {
				res = append(res, fmt.Sprintf("%s%c%s.%s", dom[:i], b, dom[i+1:], tld))
			}
		}
	}

	return helpers.Dedup(res), nil
}

func init() {
	BitSquatting = &bitsquattingStrategy{}
}
