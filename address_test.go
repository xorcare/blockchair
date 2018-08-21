// Copyright 2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchair

import (
	"encoding/json"
	"testing"
)

func TestGetAddress(t *testing.T) {
	tests := []struct {
		currency Currency
		address  string
	}{
		{Bitcoin, "1J6LK236KYiehinYCwGoSYzma9HZAgcBWP"},
		{Bitcoin, "1LUYb9y4ZUk6sox9qWtXVvNEABchKfSMir"},
		{Bitcoin, "1EHNa6Q4Jz2uvNExL497mE43ikXhwF6kZm"},
		{Bitcoin, "1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH"},
		{Bitcoin, "16rCmCmbuWDhPjWTrpQGaU3EPdZF7MTdUk"},
		{Bitcoin, "3D2oetdNuZUqQHPJmcMDDHYoqkyNVsFk9r"},
		{Bitcoin, "3Cbq7aT1tY8kMxWLbitaG7yT6bPbKChq64"},
		{Bitcoin, "3Nxwenay9Z8Lc9JBiywExpnEFiLp6Afp8v"},
		{Bitcoin, "1M2Ni8b1cW6GD7jVZLjzG6Tzs46KvzVsh7"},
		{BitcoinCash, "1J6LK236KYiehinYCwGoSYzma9HZAgcBWP"},
		{BitcoinCash, "1LUYb9y4ZUk6sox9qWtXVvNEABchKfSMir"},
		{BitcoinCash, "1EHNa6Q4Jz2uvNExL497mE43ikXhwF6kZm"},
		{BitcoinCash, "1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH"},
		{BitcoinCash, "1FeexV6bAHb8ybZjqQMjJrcCrHGW9sb6uF"},
		{BitcoinCash, "16rCmCmbuWDhPjWTrpQGaU3EPdZF7MTdUk"},
		{BitcoinCash, "1M2Ni8b1cW6GD7jVZLjzG6Tzs46KvzVsh7"},
	}
	for _, test := range tests {
		t.Run(test.currency.String(), func(t *testing.T) {
			cl := New(test.currency)
			_, e := cl.GetAddress(test.address)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestValidateBitcoinAddress(t *testing.T) {
	t.Parallel()
	tests := []struct {
		address string
		result  bool
	}{
		// one signature addresses
		{"1111111111111111111114oLvT2", true},
		{"1dyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", true},
		{"15yN7NPEpu82sHhB6TzCW5z5aXoamiKeGy", true},
		{"1F3EpcBBjVGaUuEJff9xYZBfBBALm1yfsd", true},
		// multi signature addresses
		{"3B8SEgcT9JDVKUZvm8HoKX5Av3nnn7pHqa", true},
		{"3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", true},
		// bad addresses
		{"", false},
		{"1111111111111111111114oLvT", false},
		{"1111111111111111111114iLvT", false},
		{"0111111111111111111114oLvT2", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
	}
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if ValidateBitcoinAddress(test.address) != test.result {
				t.Fatalf("validate test failed address: %s", test.address)
			}
		})
	}
}

func BenchmarkGetAddressUnmarshal(b *testing.B) {
	cl := New(Bitcoin)
	response, e := cl.GetAddress("3D2oetdNuZUqQHPJmcMDDHYoqkyNVsFk9r")
	if e != nil {
		b.Fatal(e)
	}

	bytes, e := json.Marshal(response)
	if e != nil {
		b.Fatal(e)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := json.Unmarshal(bytes, response)
		if e != nil {
			b.Fatal(e)
		}
	}
	b.StopTimer()
}

func BenchmarkGetAddressRawUnmarshal(b *testing.B) {
	cl := New(Bitcoin)
	response, e := cl.GetAddressRaw("3D2oetdNuZUqQHPJmcMDDHYoqkyNVsFk9r")
	if e != nil {
		b.Fatal(e)
	}

	bytes, e := json.Marshal(response)
	if e != nil {
		b.Fatal(e)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := json.Unmarshal(bytes, response)
		if e != nil {
			b.Fatal(e)
		}
	}
	b.StopTimer()
}

func BenchmarkGetAddressRequest(b *testing.B) {
	cl := New(Bitcoin)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, e := cl.GetAddressRaw("3D2oetdNuZUqQHPJmcMDDHYoqkyNVsFk9r")
		if e != nil {
			b.Error(e)
		}
	}
	b.StopTimer()
}
