package strategy

// Strategy defines domain generation strategy contracts
type Strategy interface {
	Generate(string) ([]string, error)
	GetName() string
}
