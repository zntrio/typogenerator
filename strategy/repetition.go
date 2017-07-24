package strategy

import (
	"fmt"

	"github.com/Zenithar/typogenerator/helpers"
)

var (
	Repetition Strategy
)

type repetitionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *repetitionStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	for i, c := range domain {
		if helpers.IsAlpha(c) {
			fuzzed := fmt.Sprintf("%s%c%c%s", domain[:i], domain[i], domain[i], domain[i+1:])
			fuzzed = combineTLD(fuzzed, tld)
			res = append(res, fuzzed)
		}
	}

	return res, nil
}

func (s *repetitionStrategy) GetName() string {
	return "Repetition"
}

func init() {
	Repetition = &repetitionStrategy{}
}
