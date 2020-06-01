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

type replaceStrategy struct {
	_mapping mapping.Mapping
}

// Replace returns a replace generation strategy
func Replace(m mapping.Mapping) Strategy {
	return &replaceStrategy{
		_mapping: m,
	}
}

// -----------------------------------------------------------------------------

func (s *replaceStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	dom := []rune(domain)

	for i := 0; i < len(dom)-1; i++ {
		keys := s._mapping.GetMapping(dom[i])
		if len(keys) > 0 {
			for _, c := range keys {
				fuzzed := fmt.Sprintf("%s%c%s", string(dom[:i]), c, string(dom[i+1:]))
				fuzzed = combineTLD(fuzzed, tld)
				res = append(res, fuzzed)
			}
		}
	}

	return helpers.Dedup(res), nil
}

func (s *replaceStrategy) GetName() string {
	return fmt.Sprintf("Replace [%s]", s._mapping.GetName())
}
