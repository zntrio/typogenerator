package strategy

import (
	"fmt"
	"strings"
)

var (
	SubDomain Strategy
)

type subdomainStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *subdomainStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := []rune(parts[0])
	tld := []rune(parts[1])

	for i := 1; i < len(dom); i++ {
		if (rune(dom[i]) != '-' || rune(dom[i]) != '.') && (rune(dom[i-1]) != '-' || rune(dom[i-1]) != '.') {
			res = append(res, fmt.Sprintf("%s.%s.%s", string(dom[:i]), string(dom[i:]), string(tld)))
		}
	}

	return res, nil
}

func init() {
	SubDomain = &subdomainStrategy{}
}
