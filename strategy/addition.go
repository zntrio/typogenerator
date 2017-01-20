package strategy

import (
	"fmt"
	"strings"
)

var (
	Addition Strategy
)

type additionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *additionStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := parts[0]
	tld := parts[1]

	for i := 97; i < 123; i++ {
		res = append(res, fmt.Sprintf("%s%c.%s", dom, i, tld))
	}

	return res, nil
}

func init() {
	Addition = &additionStrategy{}
}
