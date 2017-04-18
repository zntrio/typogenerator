package strategy

import (
	"fmt"

	"zenithar.org/go/typogenerator/helpers"
)

var (
	Repetition Strategy
)

type repetitionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *repetitionStrategy) Generate(dom string) ([]string, error) {
	res := []string{}

	for i, c := range dom {
		if helpers.IsAlpha(c) {
			res = append(res, fmt.Sprintf("%s%c%c%s", dom[:i], dom[i], dom[i], dom[i+1:]))
		}
	}

	return res, nil
}

func (s *repetitionStrategy) GetName() string {
	return "Repetition"
}

func init() {
	Repetition = &repetitionStrategy{}
}
