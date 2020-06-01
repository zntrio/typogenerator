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

package strategy

import (
	"fmt"

	"zntr.io/typogenerator/helpers"
	"zntr.io/typogenerator/mapping"
)

type similarStrategy struct {
	_mapping mapping.Mapping
}

// Similar returns a similar generation strategy
func Similar(m mapping.Mapping) Strategy {
	return &similarStrategy{
		_mapping: m,
	}
}

// -----------------------------------------------------------------------------

func (s *similarStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	dom := []rune(domain)

	for ws := range dom {
		for i := 0; i < ((len(dom) - ws) + 1); i++ {
			win := dom[i : i+ws]

			j := 0
			for j < ws {
				c := rune(win[j])

				repList := s._mapping.GetSimilar(c)
				if len(repList) > 0 {
					for _, g := range repList {
						win = []rune(fmt.Sprintf("%s%c%s", string(win[:j]), g, string(win[j+1:])))
						fuzzed := fmt.Sprintf("%s%s%s", string(dom[:i]), string(win), string(dom[i+ws:]))
						fuzzed = combineTLD(fuzzed, tld)
						res = append(res, fuzzed)
					}
				}

				j++
			}
		}
	}

	return helpers.Dedup(res), nil
}

func (s *similarStrategy) GetName() string {
	return fmt.Sprintf("Similar [%s]", s._mapping.GetName())
}
