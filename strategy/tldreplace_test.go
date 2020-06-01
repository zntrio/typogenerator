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

package strategy_test

import (
	"testing"

	"zntr.io/typogenerator/strategy"
)

func TestTLDReplace(t *testing.T) {
	out, err := strategy.TLDReplace.Generate("example", "com")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occur !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 83 {
		t.FailNow()
	}
}

func TestTLDReplaceMissingTLD(t *testing.T) {
	out, err := strategy.TLDReplace.Generate("example", "")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occur !", err)
	}

	if len(out) != 0 {
		t.FailNow()
	}
}
