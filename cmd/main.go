package main

import (
	"fmt"
	"log"
	"os"

	"zenithar.org/go/typogenerator"
	"zenithar.org/go/typogenerator/mapping"
	"zenithar.org/go/typogenerator/strategy"
)

func main() {
	all := []strategy.Strategy{
		strategy.BitSquatting,
		strategy.Homoglyph,
		strategy.Omission,
		strategy.Repetition,
		strategy.Transposition,
		strategy.Prefix,
		strategy.Hyphenation,
		strategy.DoubleHit(mapping.French),
		strategy.Similar(mapping.French),
	}

	domains, err := typogenerator.Fuzz(os.Args[1], all...)
	if err != nil {
		log.Fatal("Unable to generate domains.")
	}

	for _, domain := range domains {
		fmt.Println(domain)
	}
}
