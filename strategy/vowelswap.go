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

import "fmt"

var VowelSwap Strategy

type vowelwapStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *vowelwapStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y'}

	dom := []rune(domain)

	for i := 0; i < len(dom); i++ {
		for _, v := range vowels {
			switch dom[i] {
			case 'a', 'e', 'i', 'o', 'u', 'y':
				if dom[i] != v {
					fuzzed := fmt.Sprintf("%s%c%s", string(dom[:i]), v, string(dom[i+1:]))
					fuzzed = combineTLD(fuzzed, tld)
					res = append(res, fuzzed)
				}
			default:
			}
		}
	}

	return res, nil
}

func (s *vowelwapStrategy) GetName() string {
	return "VowelSwap"
}

func init() {
	VowelSwap = &vowelwapStrategy{}
}
