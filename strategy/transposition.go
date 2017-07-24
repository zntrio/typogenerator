package strategy

import "fmt"

var (
	Transposition Strategy
)

type transpositionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *transpositionStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	for i := 0; i < len(domain)-1; i++ {
		if domain[i+1] != domain[i] {
			fuzzed := fmt.Sprintf("%s%c%c%s", domain[:i], domain[i+1], domain[i], domain[i+2:])
			fuzzed = combineTLD(fuzzed, tld)
			res = append(res, fuzzed)
		}
	}

	return res, nil
}

func (s *transpositionStrategy) GetName() string {
	return "Transposition"
}

func init() {
	Transposition = &transpositionStrategy{}
}
