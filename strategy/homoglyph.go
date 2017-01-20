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
		'a': {"à", "á", "â", "ã", "ä", "å", "ɑ", "а"},
		'b': {"d", "lb", "ib", "ʙ", "Ь", "ｂ", "ß"},
		'c': {"ϲ", "с", "ⅽ"},
		'd': {"b", "cl", "dl", "di", "ԁ", "ժ", "ⅾ", "ｄ"},
		'e': {"é", "ê", "ë", "ē", "ĕ", "ė", "ｅ", "е"},
		'f': {"Ϝ", "Ｆ", "ｆ"},
		'g': {"q", "ɢ", "ɡ", "Ԍ", "Ԍ", "ｇ"},
		'h': {"lh", "ih", "һ", "ｈ"},
		'i': {"1", "l", "Ꭵ", "ⅰ", "ｉ"},
		'j': {"ј", "ｊ"},
		'k': {"lk", "ik", "lc", "κ", "ｋ"},
		'l': {"1", "i", "ⅼ", "ｌ"},
		'm': {"n", "nn", "rn", "rr", "ⅿ", "ｍ"},
		'n': {"m", "r", "ｎ"},
		'o': {"0", "Ο", "ο", "О", "о", "Օ", "Ｏ", "ｏ"},
		'p': {"ρ", "р", "ｐ"},
		'q': {"g", "ｑ"},
		'r': {"ʀ", "ｒ"},
		's': {"Ⴝ", "Ꮪ", "Ｓ", "ｓ"},
		't': {"τ", "ｔ"},
		'u': {"μ", "υ", "Ս", "Ｕ", "ｕ"},
		'v': {"ｖ", "ѵ", "ⅴ"},
		'w': {"vv", "ѡ", "ｗ"},
		'x': {"ⅹ", "ｘ"},
		'y': {"ʏ", "γ", "у", "Ү", "ｙ"},
		'z': {"ｚ"},
	}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := []rune(parts[0])
	tld := []rune(parts[1])

	for ws := range dom {
		for i := 0; i < ((len(dom) - ws) + 1); i++ {
			win := dom[i : i+ws]

			j := 0
			for j < ws {
				c := rune(win[j])
				if repList, ok := glyph[c]; ok {
					for _, rep := range repList {
						g := []rune(rep)

						win = []rune(fmt.Sprintf("%s%s%s", string(win[:j]), string(g), string(win[j+1:])))
						if len(g) > 1 {
							j += len(g) - 1
						}
						res = append(res, fmt.Sprintf("%s%s%s.%s", string(dom[:i]), string(win), string(dom[i+ws:]), string(tld)))
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
