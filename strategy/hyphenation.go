package strategy

import "fmt"

var (
	Hyphenation Strategy
)

type hyphenationStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *hyphenationStrategy) Generate(dom string) ([]string, error) {
	res := []string{}

	for i := 1; i < len(dom)-1; i++ {
		if (rune(dom[i]) != '-' || rune(dom[i]) != '.') && (rune(dom[i-1]) != '-' || rune(dom[i-1]) != '.') {
			res = append(res, fmt.Sprintf("%s-%s", dom[:i], dom[i:]))
		}
	}

	return res, nil
}

func (s *hyphenationStrategy) GetName() string {
	return "Hyphenation"
}

func init() {
	Hyphenation = &hyphenationStrategy{}
}
