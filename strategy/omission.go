package strategy

import (
	"fmt"
	"strings"
)

var (
	Omission Strategy
)

type omissionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *omissionStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := parts[0]
	tld := parts[1]

	for i := range dom {
		res = append(res, fmt.Sprintf("%s%s.%s", dom[:i], dom[i+1:], tld))
	}

	return res, nil
}

func init() {
	Omission = &omissionStrategy{}
}
