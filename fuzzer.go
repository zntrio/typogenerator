package typogenerator

import (
	"zenithar.org/go/typogenerator/helpers"
	"zenithar.org/go/typogenerator/strategy"
)

// Fuzz domain using given strategies
func Fuzz(domain string, strategies ...strategy.Strategy) ([]string, error) {
	res := []string{}
	var err error

	var domains []string
	for _, s := range strategies {
		if s != nil {
			domains, err = s.Generate(domain)
			if err != nil {
				break
			}

			// Copy variant to result array
			for _, variant := range domains {
				res = append(res, variant)
			}
		}
	}

	return helpers.Dedup(res), err
}
