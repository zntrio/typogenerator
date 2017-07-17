package typogenerator

import (
	"github.com/Zenithar/typogenerator/strategy"
)

// FuzzResult represents permutations results
type FuzzResult struct {
	StrategyName string   `json:"name" yaml:"name"`
	Domain       string   `json:"domain" yaml:"domain"`
	Permutations []string `json:"permutations" yaml:"permutations"`
}

// Fuzz domain using given strategies
func Fuzz(domain string, strategies ...strategy.Strategy) ([]FuzzResult, error) {
	res := []FuzzResult{}
	var err error

	var domains []string
	for _, s := range strategies {
		if s != nil {
			r := FuzzResult{
				StrategyName: s.GetName(),
				Domain:       domain,
			}

			domains, err = s.Generate(domain)
			if err != nil {
				break
			}

			// Assign permutations to result
			r.Permutations = domains

			// Add result
			res = append(res, r)
		}
	}

	return res, err
}
