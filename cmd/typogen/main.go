package main

import (
	"fmt"
	"log"

	"golang.org/x/net/idna"

	"zenithar.org/go/typogenerator"
	"zenithar.org/go/typogenerator/mapping"
	"zenithar.org/go/typogenerator/strategy"

	"github.com/namsral/flag"
)

var (
	input    = flag.String("s", "zenithar", "Defines string to alternate")
	punycode = flag.Bool("punycode", false, "Exports as punycode")
)

func init() {
	flag.Parse()
}

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

	domains, err := typogenerator.Fuzz(*input, all...)
	if err != nil {
		log.Fatal("Unable to generate domains.")
	}

	for _, domain := range domains {

		if *punycode {
			out, err := idna.ToASCII(domain)
			if err != nil {
				fmt.Println(domain)
			} else {
				fmt.Println(out)
			}
		} else {
			fmt.Println(domain)
		}

	}
}
