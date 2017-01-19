package strategy

import (
	"fmt"
	"strings"

	"zenithar.org/go/typogenerator/helpers"
)

var (
	Repetition Strategy
)

type repetitionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *repetitionStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := parts[0]
	tld := parts[1]

	for i, c := range dom {
		if helpers.IsAlpha(c) {
			res = append(res, fmt.Sprintf("%s%c%c%s.%s", dom[:i], dom[i], dom[i], dom[i+1:], tld))
		}
	}

	return res, nil
}

func init() {
	Repetition = &repetitionStrategy{}
}
