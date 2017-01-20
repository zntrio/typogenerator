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
		strategy.Addition,
		strategy.BitSquatting,
		strategy.Homoglyph,
		strategy.Omission,
		strategy.Repetition,
		strategy.Transposition,
		strategy.Prefix,
		strategy.Hyphenation,
		strategy.SubDomain,
		strategy.VowelSwap,
		strategy.Replace(mapping.French),
		strategy.DoubleHit(mapping.French),
		strategy.Similar(mapping.French),
		strategy.Replace(mapping.English),
		strategy.DoubleHit(mapping.English),
		strategy.Replace(mapping.Spanish),
		strategy.DoubleHit(mapping.Spanish),
		strategy.Similar(mapping.Spanish),
		strategy.Replace(mapping.German),
		strategy.DoubleHit(mapping.German),
		strategy.Similar(mapping.German),
	}

	domains, err := typogenerator.Fuzz(os.Args[1], all...)
	if err != nil {
		log.Fatal("Unable to generate domains.")
	}

	for _, domain := range domains {
		fmt.Println(domain)
	}
}
