package strategy

import "fmt"

var (
	TLDReplace Strategy
)

type tldReplaceStrategy struct {
}

// Top list generated from:
// https://w3techs.com/technologies/overview/top_level_domain/all
var topTLDs = []string{
	"com",
	"ru",
	"org",
	"net",
	"de",
	"jp",
	"uk",
	"br",
	"it",
	"pl",
	"fr",
	"in",
	"ir",
	"au",
	"info",
	"cn",
	"nl",
	"es",
	"cz",
	"kr",
	"ca",
	"ua",
	"eu",
	"co",
	"gr",
	"ro",
	"za",
	"ch",
	"se",
	"tw",
	"biz",
	"hu",
	"vn",
	"mx",
	"be",
	"at",
	"tr",
	"dk",
	"me",
	"ar",
	"tv",
	"sk",
	"no",
	"us",
	"fi",
	"cl",
	"id",
	"io",
	"xyz",
	"pt",
	"by",
	"il",
	"ie",
	"nz",
	"kz",
	"lt",
	"hk",
	"cc",
	"my",
	"club",
	"sg",
	"top",
	"bg",
	"рф",
	"edu",
	"th",
	"su",
	"pk",
	"hr",
	"rs",
	"pro",
	"si",
	"lv",
	"az",
	"pe",
	"download",
	"ae",
	"ph",
	"ee",
	"online",
	"ng",
	"pw",
	"cat",
	"ve",
}

// -----------------------------------------------------------------------------

func (s *tldReplaceStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	// Skip strategy if input is missing TLD
	if tld == "" {
		return res, nil
	}

	for _, topTLD := range topTLDs {
		if tld == topTLD {
			continue
		}
		fuzzed := fmt.Sprintf("%s.%s", domain, topTLD)
		res = append(res, fuzzed)
	}

	return res, nil
}

func init() {
	TLDReplace = &tldReplaceStrategy{}
}

func (s *tldReplaceStrategy) GetName() string {
	return "TLDReplace"
}
