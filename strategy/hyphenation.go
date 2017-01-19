package strategy

import (
	"fmt"
	"strings"
)

var (
	Hyphenation Strategy
)

type hyphenationStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *hyphenationStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := parts[0]
	tld := parts[1]

	for i := 1; i < len(dom)-1; i++ {
		if (rune(dom[i]) != '-' || rune(dom[i]) != '.') && (rune(dom[i-1]) != '-' || rune(dom[i-1]) != '.') {
			res = append(res, fmt.Sprintf("%s-%s.%s", dom[:i], dom[i:], tld))
		}
	}

	return res, nil
}

func init() {
	Hyphenation = &hyphenationStrategy{}
}
