package strategy_test

import (
	"testing"

	"go.zenithar.org/typogenerator/mapping"
	"go.zenithar.org/typogenerator/strategy"
)

func TestReplace(t *testing.T) {
	out, err := strategy.Replace(mapping.French).Generate("zenithar", "")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 41 {
		t.FailNow()
	}
}
