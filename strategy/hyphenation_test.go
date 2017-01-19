package strategy_test

import (
	"testing"

	"zenithar.org/go/typogenerator/strategy"
)

func TestHyphenation(t *testing.T) {
	out, err := strategy.Hyphenation.Generate("zemithar.org")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 6 {
		t.FailNow()
	}
}
