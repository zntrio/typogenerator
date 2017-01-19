package strategy

import (
	"fmt"
	"strings"

	"zenithar.org/go/typogenerator/helpers"
	"zenithar.org/go/typogenerator/mapping"
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

func (s *doublehitStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := []rune(parts[0])
	tld := []rune(parts[1])

	for i := 0; i < len(dom)-1; i++ {
		keys := s._mapping.GetMapping(dom[i])
		if len(keys) > 0 {
			for _, c := range keys {
				res = append(res, fmt.Sprintf("%s%c%c%s.%s", string(dom[:i]), c, dom[i], string(dom[i+1:]), string(tld)))
				res = append(res, fmt.Sprintf("%s%c%c%s.%s", string(dom[:i]), dom[i], c, string(dom[i+1:]), string(tld)))
			}
		}
	}

	return helpers.Dedup(res), nil
}
