package strategy

import "fmt"

var (
	Hyphenation Strategy
)

type hyphenationStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *hyphenationStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	for i := 1; i < len(domain)-1; i++ {
		if (rune(domain[i]) != '-' || rune(domain[i]) != '.') && (rune(domain[i-1]) != '-' || rune(domain[i-1]) != '.') {
			fuzzed := fmt.Sprintf("%s-%s", domain[:i], domain[i:])
			fuzzed = combineTLD(fuzzed, tld)
			res = append(res, fuzzed)
		}
	}

	return res, nil
}

func (s *hyphenationStrategy) GetName() string {
	return "Hyphenation"
}

func init() {
	Hyphenation = &hyphenationStrategy{}
}
