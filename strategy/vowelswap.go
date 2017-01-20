package strategy

import (
	"fmt"
	"strings"
)

var (
	VowelSwap Strategy
)

type vowelwapStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *vowelwapStrategy) Generate(domain string) ([]string, error) {
	res := []string{}
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y'}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := []rune(parts[0])
	tld := []rune(parts[1])

	for i := 0; i < len(dom); i++ {
		for _, v := range vowels {
			switch dom[i] {
			case 'a', 'e', 'i', 'o', 'u', 'y':
				res = append(res, fmt.Sprintf("%s%c%s.%s", string(dom[:i]), v, string(dom[i+1:]), string(tld)))
			default:
			}
		}
	}

	return res, nil
}

func init() {
	VowelSwap = &vowelwapStrategy{}
}
