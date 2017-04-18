package strategy

import "fmt"

var (
	Transposition Strategy
)

type transpositionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *transpositionStrategy) Generate(dom string) ([]string, error) {
	res := []string{}

	for i := 0; i < len(dom)-1; i++ {
		if dom[i+1] != dom[i] {
			res = append(res, fmt.Sprintf("%s%c%c%s", dom[:i], dom[i+1], dom[i], dom[i+2:]))
		}
	}

	return res, nil
}

func (s *transpositionStrategy) GetName() string {
	return "Transposition"
}

func init() {
	Transposition = &transpositionStrategy{}
}
