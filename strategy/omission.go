package strategy

import "fmt"

var (
	Omission Strategy
)

type omissionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *omissionStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	for i := range domain {
		fuzzed := fmt.Sprintf("%s%s", domain[:i], domain[i+1:])
		fuzzed = combineTLD(fuzzed, tld)
		res = append(res, fuzzed)
	}

	return res, nil
}

func (s *omissionStrategy) GetName() string {
	return "Omission"
}

func init() {
	Omission = &omissionStrategy{}
}
