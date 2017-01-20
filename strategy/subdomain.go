package strategy

import "fmt"

var (
	SubDomain Strategy
)

type subdomainStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *subdomainStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	dom := []rune(domain)

	for i := 1; i < len(dom); i++ {
		if (rune(dom[i]) != '-' || rune(dom[i]) != '.') && (rune(dom[i-1]) != '-' || rune(dom[i-1]) != '.') {
			res = append(res, fmt.Sprintf("%s.%s", string(dom[:i]), string(dom[i:])))
		}
	}

	return res, nil
}

func init() {
	SubDomain = &subdomainStrategy{}
}
