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

var Repetition Strategy

type repetitionStrategy struct {
}

// -----------------------------------------------------------------------------

func (s *repetitionStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	for i, c := range domain {
		if helpers.IsAlpha(c) {
			fuzzed := fmt.Sprintf("%s%c%c%s", domain[:i], domain[i], domain[i], domain[i+1:])
			fuzzed = combineTLD(fuzzed, tld)
			res = append(res, fuzzed)
		}
	}

	return res, nil
}

func (s *repetitionStrategy) GetName() string {
	return "Repetition"
}

func init() {
	Repetition = &repetitionStrategy{}
}
