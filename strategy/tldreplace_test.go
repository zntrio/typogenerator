package strategy_test

import (
	"testing"

	"go.zenithar.org/typogenerator/strategy"
)

func TestTLDReplace(t *testing.T) {
	out, err := strategy.TLDReplace.Generate("example", "com")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occur !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 83 {
		t.FailNow()
	}
}

func TestTLDReplaceMissingTLD(t *testing.T) {
	out, err := strategy.TLDReplace.Generate("example", "")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occur !", err)
	}

	if len(out) != 0 {
		t.FailNow()
	}
}
