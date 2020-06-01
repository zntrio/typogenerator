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

var TLDReplace Strategy

type tldReplaceStrategy struct {
}

// Top list generated from:
// https://w3techs.com/technologies/overview/top_level_domain/all
var topTLDs = []string{
	"com",
	"ru",
	"org",
	"net",
	"de",
	"jp",
	"uk",
	"br",
	"it",
	"pl",
	"fr",
	"in",
	"ir",
	"au",
	"info",
	"cn",
	"nl",
	"es",
	"cz",
	"kr",
	"ca",
	"ua",
	"eu",
	"co",
	"gr",
	"ro",
	"za",
	"ch",
	"se",
	"tw",
	"biz",
	"hu",
	"vn",
	"mx",
	"be",
	"at",
	"tr",
	"dk",
	"me",
	"ar",
	"tv",
	"sk",
	"no",
	"us",
	"fi",
	"cl",
	"id",
	"io",
	"xyz",
	"pt",
	"by",
	"il",
	"ie",
	"nz",
	"kz",
	"lt",
	"hk",
	"cc",
	"my",
	"club",
	"sg",
	"top",
	"bg",
	"рф",
	"edu",
	"th",
	"su",
	"pk",
	"hr",
	"rs",
	"pro",
	"si",
	"lv",
	"az",
	"pe",
	"download",
	"ae",
	"ph",
	"ee",
	"online",
	"ng",
	"pw",
	"cat",
	"ve",
}

// -----------------------------------------------------------------------------

func (s *tldReplaceStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	// Skip strategy if input is missing TLD
	if tld == "" {
		return res, nil
	}

	for _, topTLD := range topTLDs {
		if tld == topTLD {
			continue
		}
		fuzzed := fmt.Sprintf("%s.%s", domain, topTLD)
		res = append(res, fuzzed)
	}

	return res, nil
}

func init() {
	TLDReplace = &tldReplaceStrategy{}
}

func (s *tldReplaceStrategy) GetName() string {
	return "TLDReplace"
}
