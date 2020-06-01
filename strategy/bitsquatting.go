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
)

var BitSquatting Strategy

type bitsquattingStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *bitsquattingStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	masks := []int32{1, 2, 4, 8, 16, 32, 64, 128}

	for i, c := range domain {
		for m := range masks {
			b := rune(int(c) ^ m)
			o := int(b)
			if (o >= 48 && o <= 57) || (o >= 97 && o <= 122) || o == 45 {
				fuzzed := fmt.Sprintf("%s%c%s", domain[:i], b, domain[i+1:])
				fuzzed = combineTLD(fuzzed, tld)
				res = append(res, fuzzed)
			}
		}
	}

	return helpers.Dedup(res), nil
}

func (s *bitsquattingStrategy) GetName() string {
	return "BitSquatting"
}

func init() {
	BitSquatting = &bitsquattingStrategy{}
}
