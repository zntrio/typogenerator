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

package mapping

// Mapping defines mapping contract
type Mapping interface {
	GetMapping(rune) []rune
	GetSimilar(rune) []rune
	GetName() string
}

type defaultMapping struct {
	keyboard map[rune][]rune
	similar  map[rune][]rune
	name     string
}

// -----------------------------------------------------------------------------

func (m *defaultMapping) GetName() string {
	return m.name
}

func (m *defaultMapping) GetMapping(r rune) []rune {
	if values, ok := m.keyboard[r]; ok {
		return values
	}
	return []rune{}
}

func (m *defaultMapping) GetSimilar(r rune) []rune {
	if values, ok := m.similar[r]; ok {
		return values
	}
	return []rune{}
}
