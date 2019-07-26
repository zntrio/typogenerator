package strategy

import (
	"fmt"

	"go.zenithar.org/typogenerator/helpers"
	"go.zenithar.org/typogenerator/mapping"
)

type doublehitStrategy struct {
	_mapping mapping.Mapping
}

// DoubleHit returns a double hit genration strategy
func DoubleHit(m mapping.Mapping) Strategy {
	return &doublehitStrategy{
		_mapping: m,
	}
}

// -----------------------------------------------------------------------------

func (s *doublehitStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	dom := []rune(domain)

	for i := 0; i < len(dom)-1; i++ {
		keys := s._mapping.GetMapping(dom[i])
		if len(keys) > 0 {
			for _, c := range keys {
				fuzzed := fmt.Sprintf("%s%c%c%s", string(dom[:i]), c, dom[i], string(dom[i+1:]))
				fuzzed = combineTLD(fuzzed, tld)
				res = append(res, fuzzed)

				fuzzed = fmt.Sprintf("%s%c%c%s", string(dom[:i]), dom[i], c, string(dom[i+1:]))
				fuzzed = combineTLD(fuzzed, tld)
				res = append(res, fuzzed)
			}
		}
	}

	return helpers.Dedup(res), nil
}

func (s *doublehitStrategy) GetName() string {
	return fmt.Sprintf("DoubleHit [%s]", s._mapping.GetName())
}
