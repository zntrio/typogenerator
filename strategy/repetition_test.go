package strategy_test

import (
	"testing"

	"go.zenithar.org/typogenerator/strategy"
)

func TestRepetition(t *testing.T) {
	out, err := strategy.Repetition.Generate("zemithar", "")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 8 {
		t.FailNow()
	}
}
