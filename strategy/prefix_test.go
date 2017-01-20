package strategy_test

import (
	"testing"

	"zenithar.org/go/typogenerator/strategy"
)

func TestPrefix(t *testing.T) {
	out, err := strategy.Prefix.Generate("zemithar.org")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 14 {
		t.FailNow()
	}
}
