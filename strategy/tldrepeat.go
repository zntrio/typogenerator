package strategy

import "fmt"

var (
	TLDRepeat Strategy
)

type tldRepeatStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *tldRepeatStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	// Skip strategy if input is missing TLD
	if tld == "" {
		return res, nil
	}

	fuzzed := fmt.Sprintf("%s%s.%s", domain, tld, tld)
	res = append(res, fuzzed)

	return res, nil
}

func init() {
	TLDRepeat = &tldRepeatStrategy{}
}

func (s *tldRepeatStrategy) GetName() string {
	return "TLDRepeat"
}
