package strategy

func combineTLD(domain, tld string) string {
	if tld == "" {
		return domain
	}

	return domain + "." + tld
}
