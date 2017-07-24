package typogenerator

import (
	"github.com/Zenithar/typogenerator/strategy"
	"github.com/weppos/publicsuffix-go/publicsuffix"
)

// FuzzResult represents permutations results
type FuzzResult struct {
	StrategyName string   `json:"name" yaml:"name"`
	Domain       string   `json:"domain" yaml:"domain"`
	Permutations []string `json:"permutations" yaml:"permutations"`
}

// Fuzz string using given strategies
func Fuzz(name string, strategies ...strategy.Strategy) ([]FuzzResult, error) {
	return fuzz(name, "", strategies...)
}

// FuzzDomain splits a domain into (TRD + SLD) + TLD and fuzzes using given strategies
func FuzzDomain(domain string, strategies ...strategy.Strategy) ([]FuzzResult, error) {
	parsed, err := publicsuffix.Parse(domain)
	if err != nil {
		return []FuzzResult{}, err
	}

	domain = parsed.SLD
	if parsed.TRD != "" {
		domain = parsed.TRD + "." + domain
	}

	return fuzz(domain, parsed.TLD, strategies...)
}

func fuzz(domain string, tld string, strategies ...strategy.Strategy) ([]FuzzResult, error) {
	res := []FuzzResult{}
	var err error

	var domains []string
	for _, s := range strategies {
		if s != nil {
			r := FuzzResult{
				StrategyName: s.GetName(),
				Domain:       domain,
			}

			domains, err = s.Generate(domain, tld)
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
