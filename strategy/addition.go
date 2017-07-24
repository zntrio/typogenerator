package strategy

import "fmt"

var (
	Addition Strategy
)

type additionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *additionStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	for i := 97; i < 123; i++ {
		fuzzed := fmt.Sprintf("%s%c", domain, i)
		fuzzed = combineTLD(fuzzed, tld)
		res = append(res, fuzzed)
	}

	return res, nil
}

func init() {
	Addition = &additionStrategy{}
}

func (s *additionStrategy) GetName() string {
	return "Addition"
}
