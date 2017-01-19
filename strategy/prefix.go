package strategy

import "fmt"

var (
	Prefix Strategy
)

type prefixStrategy struct {
	prefixes []string
}

// -----------------------------------------------------------------------------

func (s *prefixStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	for _, prefix := range s.prefixes {
		res = append(res, fmt.Sprintf("%s%s", prefix, domain))
	}

	return res, nil
}

func init() {
	Prefix = &prefixStrategy{
		prefixes: []string{
			"www",
			"www-",
		},
	}
}
