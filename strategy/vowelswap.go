package strategy

import "fmt"

var (
	VowelSwap Strategy
)

type vowelwapStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *vowelwapStrategy) Generate(domain string) ([]string, error) {
	res := []string{}
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y'}

	dom := []rune(domain)

	for i := 0; i < len(dom); i++ {
		for _, v := range vowels {
			switch dom[i] {
			case 'a', 'e', 'i', 'o', 'u', 'y':
				res = append(res, fmt.Sprintf("%s%c%s", string(dom[:i]), v, string(dom[i+1:])))
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
