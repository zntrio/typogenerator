package strategy_test

import (
	"testing"

	"go.zenithar.org/typogenerator/strategy"
)

func TestHomoglyph(t *testing.T) {
	out, err := strategy.Homoglyph.Generate("zenithar", "")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 90 {
		t.FailNow()
	}
}
