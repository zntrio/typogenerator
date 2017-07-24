package strategy

// Strategy defines domain generation strategy contracts
type Strategy interface {
	Generate(domain, tld string) ([]string, error)
	GetName() string
}
