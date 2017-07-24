package strategy

import "fmt"

var (
	VowelSwap Strategy
)

type vowelwapStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *vowelwapStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y'}

	dom := []rune(domain)

	for i := 0; i < len(dom); i++ {
		for _, v := range vowels {
			switch dom[i] {
			case 'a', 'e', 'i', 'o', 'u', 'y':
				if dom[i] != v {
					fuzzed := fmt.Sprintf("%s%c%s", string(dom[:i]), v, string(dom[i+1:]))
					fuzzed = combineTLD(fuzzed, tld)
					res = append(res, fuzzed)
				}
			default:
			}
		}
	}

	return res, nil
}

func (s *vowelwapStrategy) GetName() string {
	return "VowelSwap"
}

func init() {
	VowelSwap = &vowelwapStrategy{}
}
