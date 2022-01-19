// Licensed to Typogenerator under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Typogenerator licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hduplooy/gocsv"
	"github.com/namsral/flag"
	"golang.org/x/net/idna"

	"zntr.io/typogenerator"
	"zntr.io/typogenerator/mapping"
	"zntr.io/typogenerator/strategy"
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
		if err := writer.Write([]string{"strategy", "domain", "permunation", "idna"}); err != nil {
			panic(err)
		}

		for _, r := range results {
			for _, p := range r.Permutations {
				puny, _ := idna.ToASCII(p)
				if err := writer.Write([]string{r.StrategyName, r.Domain, p, puny}); err != nil {
					panic(err)
				}
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
