package main

import (
	"fmt"
	"log"

	"zenithar.org/go/typogenerator"
	"zenithar.org/go/typogenerator/strategy"
)

func main() {
	domains, err := typogenerator.Fuzz("zenithar.org", strategy.BitSquatting, strategy.Homoglyph, strategy.Insertion, strategy.Omission, strategy.Repetition, strategy.Transposition)
	if err != nil {
		log.Fatal("Unable to generate domains.")
	}

	for _, domain := range domains {
		fmt.Println(domain)
	}
}
