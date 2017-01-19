package strategy

import (
	"fmt"
	"strings"

	"zenithar.org/go/typogenerator/helpers"
)

var (
	Homoglyph Strategy
)

type homoglyphStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *homoglyphStrategy) Generate(domain string) ([]string, error) {
	res := []string{}
	glyph := map[rune][]string{
		'd': {"b", "cl"},
		'm': {"r", "rn"},
		'l': {"1", "i"},
		'o': {"O"},
		'w': {"vv"},
		'n': {"m"},
		'b': {"d"},
		'i': {"l"},
		'g': {"q"},
		'q': {"g"},
	}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := parts[0]
	tld := parts[1]

	for ws := range dom {
		for i := 0; i < ((len(dom) - ws) + 1); i++ {
			win := dom[i : i+ws]

			j := 0
			for j < ws {
				c := rune(win[j])
				if repList, ok := glyph[c]; ok {
					for _, g := range repList {
						win = fmt.Sprintf("%s%s%s", win[:j], g, win[j+1:])

						if len(g) > 1 {
							j += len(g) - 1
						}
						res = append(res, fmt.Sprintf("%s%s%s.%s", dom[:i], win, dom[i+ws:], tld))
					}
				}
				j++
			}
		}
	}

	return helpers.Dedup(res), nil
}

func init() {
	Homoglyph = &homoglyphStrategy{}
}
