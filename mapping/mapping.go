package mapping

// Mapping defines mapping contract
type Mapping interface {
	GetMapping(rune) []rune
	GetSimilar(rune) []rune
}

type defaultMapping struct {
	keyboard map[rune][]rune
	similar  map[rune][]rune
}

// -----------------------------------------------------------------------------

func (m *defaultMapping) GetMapping(r rune) []rune {
	if values, ok := m.keyboard[r]; ok {
		return values
	}
	return []rune{}
}

func (m *defaultMapping) GetSimilar(r rune) []rune {
	if values, ok := m.similar[r]; ok {
		return values
	}
	return []rune{}
}
