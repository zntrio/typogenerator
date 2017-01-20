package strategy

import "fmt"

var (
	Addition Strategy
)

type additionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *additionStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	for i := 97; i < 123; i++ {
		res = append(res, fmt.Sprintf("%s%c", domain, i))
	}

	return res, nil
}

func init() {
	Addition = &additionStrategy{}
}
