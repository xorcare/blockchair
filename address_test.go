// Copyright 2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchair

import (
	"testing"
)

func TestClient_GetAddress(t *testing.T) {
	tests := []struct {
		Name     string
		Currency uint8
		Address  string
	}{
		{currencyURL[Bitcoin], Bitcoin, "1LUYb9y4ZUk6sox9qWtXVvNEABchKfSMir"},
		{currencyURL[Bitcoin], Bitcoin, "1EHNa6Q4Jz2uvNExL497mE43ikXhwF6kZm"},
		{currencyURL[Bitcoin], Bitcoin, "1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH"},
		{currencyURL[BitcoinCash], BitcoinCash, "1LUYb9y4ZUk6sox9qWtXVvNEABchKfSMir"},
		{currencyURL[BitcoinCash], BitcoinCash, "1EHNa6Q4Jz2uvNExL497mE43ikXhwF6kZm"},
		{currencyURL[BitcoinCash], BitcoinCash, "1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH"},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			cl := New(test.Currency)
			a, e := cl.GetAddress(test.Address)
			if e != nil {
				t.Fatal(e)
			}
			t.Log(a)
		})
	}
}
