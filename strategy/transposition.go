package strategy

import (
	"fmt"
	"strings"
)

var (
	Transposition Strategy
)

type transpositionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *transpositionStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := parts[0]
	tld := parts[1]

	for i := 0; i < len(dom)-1; i++ {
		if dom[i+1] != dom[i] {
			res = append(res, fmt.Sprintf("%s%c%c%s.%s", dom[:i], dom[i+1], dom[i], dom[i+2:], tld))
		}
	}

	return res, nil
}

func init() {
	Transposition = &transpositionStrategy{}
}
