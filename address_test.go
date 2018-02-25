// Copyright 2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchair

import (
	"testing"
)

func TestClient_GetAddress(t *testing.T) {
	tests := []struct {
		name     string
		currency Currency
		address  string
	}{
		{currencyURL[Bitcoin], Bitcoin, "1J6LK236KYiehinYCwGoSYzma9HZAgcBWP"},
		{currencyURL[Bitcoin], Bitcoin, "1LUYb9y4ZUk6sox9qWtXVvNEABchKfSMir"},
		{currencyURL[Bitcoin], Bitcoin, "1EHNa6Q4Jz2uvNExL497mE43ikXhwF6kZm"},
		{currencyURL[Bitcoin], Bitcoin, "1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH"},
		{currencyURL[Bitcoin], Bitcoin, "16rCmCmbuWDhPjWTrpQGaU3EPdZF7MTdUk"},
		{currencyURL[Bitcoin], Bitcoin, "3D2oetdNuZUqQHPJmcMDDHYoqkyNVsFk9r"},
		{currencyURL[Bitcoin], Bitcoin, "3Cbq7aT1tY8kMxWLbitaG7yT6bPbKChq64"},
		{currencyURL[Bitcoin], Bitcoin, "3Nxwenay9Z8Lc9JBiywExpnEFiLp6Afp8v"},
		{currencyURL[Bitcoin], Bitcoin, "1M2Ni8b1cW6GD7jVZLjzG6Tzs46KvzVsh7"},
		{currencyURL[BitcoinCash], BitcoinCash, "1J6LK236KYiehinYCwGoSYzma9HZAgcBWP"},
		{currencyURL[BitcoinCash], BitcoinCash, "1LUYb9y4ZUk6sox9qWtXVvNEABchKfSMir"},
		{currencyURL[BitcoinCash], BitcoinCash, "1EHNa6Q4Jz2uvNExL497mE43ikXhwF6kZm"},
		{currencyURL[BitcoinCash], BitcoinCash, "1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH"},
		{currencyURL[BitcoinCash], BitcoinCash, "1FeexV6bAHb8ybZjqQMjJrcCrHGW9sb6uF"},
		{currencyURL[BitcoinCash], BitcoinCash, "16rCmCmbuWDhPjWTrpQGaU3EPdZF7MTdUk"},
		{currencyURL[BitcoinCash], BitcoinCash, "1M2Ni8b1cW6GD7jVZLjzG6Tzs46KvzVsh7"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cl := New(test.currency)
			_, e := cl.GetAddress(test.address)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}
