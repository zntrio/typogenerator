package strategy_test

import (
	"testing"

	"github.com/Zenithar/typogenerator/strategy"
)

func TestOmission(t *testing.T) {
	out, err := strategy.Omission.Generate("zemithar", "")
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
