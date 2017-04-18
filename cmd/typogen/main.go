package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/idna"

	"zenithar.org/go/typogenerator"
	"zenithar.org/go/typogenerator/mapping"
	"zenithar.org/go/typogenerator/strategy"

	"github.com/hduplooy/gocsv"
	"github.com/namsral/flag"
)

var (
	input           = flag.String("s", "zenithar", "Defines string to alternate")
	permutationOnly = flag.Bool("p", false, "Display permutted domain only")
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

	results, err := typogenerator.Fuzz(*input, all...)
	if err != nil {
		log.Fatal("Unable to generate domains.")
	}

	if !*permutationOnly {
		writer := gocsv.NewWriter(os.Stdout)
		writer.QuoteFields = true
		defer writer.Flush()

		// Write headers
		writer.Write([]string{"strategy", "domain", "permunation", "idna"})

		for _, r := range results {
			for _, p := range r.Permutations {
				puny, _ := idna.ToASCII(p)
				writer.Write([]string{r.StrategyName, r.Domain, p, puny})
			}
		}
	} else {
		for _, r := range results {
			for _, p := range r.Permutations {
				fmt.Println(p)
			}
		}
	}

}
