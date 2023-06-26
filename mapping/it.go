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

// Italian mapping
var Italian Mapping

func init() {
	Italian = &defaultMapping{
		name: "Italy",
		keyboard: map[rune][]rune{
			'q': {'1', '2', 'w', 's', 'a', 'o'},
			'w': {'1', '2', '3', 'e', 'd', 's', 'a', 'q', 'v'},
			'e': {'2', '3', '4', 'r', 'f', 'd', 's', 'w'},
			'r': {'3', '4', '5', 't', 'g', 'f', 'd', 'e'},
			't': {'4', '5', '6', 'y', 'h', 'g', 'f', 'r'},
			'y': {'5', '6', '7', 'u', 'j', 'h', 'g', 't', 'i'},
			'u': {'7', '8', 'i', 'k', 'j', 'h', 'y'},
			'i': {'8', '9', 'o', 'k', 'j', 'u', 'l', 'y'},
			'o': {'9', '0', 'p', 'l', 'k', 'i', 'q'},
			'p': {'l', 'o', '0'},
			'a': {'q', 'w', 's', 'x', 'z', 'e'},
			's': {'q', 'w', 'e', 'd', 'c', 'x', 'z', 'a'},
			'd': {'w', 'e', 'r', 'f', 'v', 'c', 'x', 's'},
			'f': {'e', 'r', 't', 'g', 'b', 'v', 'c', 'd'},
			'g': {'r', 't', 'y', 'h', 'n', 'b', 'v', 'f'},
			'h': {'t', 'y', 'u', 'j', 'm', 'n', 'b', 'g'},
			'j': {'y', 'u', 'i', 'k', 'm', 'n', 'h'},
			'k': {'u', 'i', 'o', 'l', 'm', 'j'},
			'l': {'i', 'o', 'p', 'k'},
			'z': {'x', 's', 'a', 'q', 'w'},
			'x': {'z', 'a', 's', 'd', 'c'},
			'c': {'x', 's', 'd', 'f', 'v'},
			'v': {'c', 'd', 'f', 'g', 'b', 'w'},
			'b': {'v', 'f', 'g', 'h', 'n'},
			'n': {'b', 'g', 'h', 'j', 'm'},
			'm': {'n', 'h', 'j', 'k'},
			'1': {'q', 'w', '2'},
			'2': {'1', 'q', 'w', 'e', '3'},
			'3': {'2', 'w', 'e', 'r', '4'},
			'4': {'3', 'e', 'r', 't', '5'},
			'5': {'4', 'r', 't', '6'},
			'6': {'5', 't', 'y', '7'},
			'7': {'6', 'y', 'u', '8'},
			'8': {'7', 'u', 'i', '9'},
			'9': {'8', 'i', 'o', '0'},
			'0': {'9', 'p', 'o'},
		},
		similar: map[rune][]rune{
			'à': {'a'},
			'è': {'e'},
			'é': {'e'},
			'ì': {'i'},
			'ò': {'o'},
			'ù': {'u'},
			'a': {'à'},
			'e': {'è', 'é'},
			'i': {'ì'},
			'o': {'ò'},
			'u': {'ù'},
		},
	}
}
