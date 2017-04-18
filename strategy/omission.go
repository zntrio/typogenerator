package strategy

import "fmt"

var (
	Omission Strategy
)

type omissionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *omissionStrategy) Generate(dom string) ([]string, error) {
	res := []string{}

	for i := range dom {
		res = append(res, fmt.Sprintf("%s%s", dom[:i], dom[i+1:]))
	}

	return res, nil
}

func (s *omissionStrategy) GetName() string {
	return "Omission"
}

func init() {
	Omission = &omissionStrategy{}
}
