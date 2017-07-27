package strategy_test

import (
	"testing"

	"github.com/Zenithar/typogenerator/strategy"
)

func TestTLDRepeat(t *testing.T) {
	out, err := strategy.TLDRepeat.Generate("example", "com")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occur !", err)
	}

	if len(out) != 1 {
		t.FailNow()
	}

	if out[0] != "examplecom.com" {
		t.FailNow()
	}
}

func TestTLDRepeatMissingTLD(t *testing.T) {
	out, err := strategy.TLDRepeat.Generate("example", "")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occur !", err)
	}

	if len(out) != 0 {
		t.FailNow()
	}
}
