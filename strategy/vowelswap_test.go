package strategy_test

import (
	"testing"

	"github.com/Zenithar/typogenerator/strategy"
)

func TestVowelSwap(t *testing.T) {
	in := "zenithar"
	out, err := strategy.VowelSwap.Generate(in, "")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 15 {
		t.FailNow()
	}

	for _, res := range out {
		if res == in {
			t.Fail()
			t.Fatal("Vowel swap should not swap a letter with itself!")
		}
	}
}
