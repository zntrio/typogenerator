package strategy_test

import (
	"testing"

	"go.zenithar.org/typogenerator/strategy"
)

func TestPrefix(t *testing.T) {
	out, err := strategy.Prefix.Generate("zenithar", "")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 22 {
		t.FailNow()
	}
}
