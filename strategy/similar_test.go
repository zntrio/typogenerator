package strategy_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"zenithar.org/go/typogenerator/mapping"
	"zenithar.org/go/typogenerator/strategy"
)

func TestSimilar(t *testing.T) {
	out, err := strategy.Similar(mapping.French).Generate("zenithar.org")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	spew.Dump(out)
	if len(out) != 17 {
		t.FailNow()
	}
}
