package strategy_test

import (
	"testing"

	"zenithar.org/go/typogenerator/strategy"
)

func TestHomoglyph(t *testing.T) {
	out, err := strategy.Homoglyph.Generate("zemithar")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 87 {
		t.FailNow()
	}
}
