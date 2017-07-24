package strategy

import "fmt"

var (
	Prefix Strategy
)

type prefixStrategy struct {
	prefixes []string
}

// -----------------------------------------------------------------------------

func (s *prefixStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	for _, prefix := range s.prefixes {
		fuzzed := fmt.Sprintf("%s%s", prefix, domain)
		fuzzed = combineTLD(fuzzed, tld)
		res = append(res, fuzzed)
	}

	return res, nil
}

func (s *prefixStrategy) GetName() string {
	return "Prefix"
}

func init() {
	Prefix = &prefixStrategy{
		prefixes: []string{
			"www",
			"www-",
			"mail",
			"mail-",
			"login",
			"login-",
			"auth",
			"auth-",
			"ftp",
			"ftp-",
			"smtp",
			"smtp-",
			"imap",
			"imap-",
			"account",
			"account-",
			"accounts",
			"accounts-",
			"users",
			"users-",
			"sso",
			"sso-",
		},
	}
}
