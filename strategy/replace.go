package strategy

import (
	"fmt"

	"github.com/Zenithar/typogenerator/helpers"
	"github.com/Zenithar/typogenerator/mapping"
)

type replaceStrategy struct {
	_mapping mapping.Mapping
}

// Replace returns a replace generation strategy
func Replace(m mapping.Mapping) Strategy {
	return &replaceStrategy{
		_mapping: m,
	}
}

// -----------------------------------------------------------------------------

func (s *replaceStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	dom := []rune(domain)

	for i := 0; i < len(dom)-1; i++ {
		keys := s._mapping.GetMapping(dom[i])
		if len(keys) > 0 {
			for _, c := range keys {
				fuzzed := fmt.Sprintf("%s%c%s", string(dom[:i]), c, string(dom[i+1:]))
				fuzzed = combineTLD(fuzzed, tld)
				res = append(res, fuzzed)
			}
		}
	}

	return helpers.Dedup(res), nil
}

func (s *replaceStrategy) GetName() string {
	return fmt.Sprintf("Replace [%s]", s._mapping.GetName())
}
