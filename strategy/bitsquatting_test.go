package strategy_test

import (
	"testing"

	"zenithar.org/go/typogenerator/strategy"
)

func TestBitSquatting(t *testing.T) {
	out, err := strategy.BitSquatting.Generate("zenithar")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 50 {
		t.FailNow()
	}
}
