package strategy

import "fmt"

var (
	SubDomain Strategy
)

type subdomainStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *subdomainStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	dom := []rune(domain)

	for i := 1; i < len(dom); i++ {
		if !(rune(dom[i]) == '-' || rune(dom[i]) == '.') && !(rune(dom[i-1]) == '-' || rune(dom[i-1]) == '.') {
			fuzzed := fmt.Sprintf("%s.%s", string(dom[:i]), string(dom[i:]))
			fuzzed = combineTLD(fuzzed, tld)
			res = append(res, fuzzed)
		}
	}

	return res, nil
}

func (s *subdomainStrategy) GetName() string {
	return "Subdomain"
}

func init() {
	SubDomain = &subdomainStrategy{}
}
