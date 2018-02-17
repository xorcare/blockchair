// Copyright 2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchair

import (
	"fmt"
)

type AddressResponse struct {
	Data   []Address `json:"data"`
	Rows   uint      `json:"rows"`
	Limit  int64     `json:"limit"`
	Time   float64   `json:"time"`
	Cache  int       `json:"cache"`
	Source string    `json:"source"`
}

type Address struct {
	SumValue            string     `json:"sum_value"`
	SumValueUsd         float64    `json:"sum_value_usd"`
	SumSpendingValueUsd float64    `json:"sum_spending_value_usd"`
	MaxTimeReceiving    string     `json:"max_time_receiving"`
	MaxTimeSpending     string     `json:"max_time_spending"`
	MinTimeReceiving    string     `json:"min_time_receiving"`
	CountTotal          string     `json:"count_total"`
	Rate                string     `json:"rate"`
	SumValueUnspent     int64      `json:"sum_value_unspent"`
	SumValueUnspentUsd  int64      `json:"sum_value_unspent_usd"`
	CountUnspent        int64      `json:"count_unspent"`
	PluUsd              int64      `json:"plu_usd"`
	MinTimeSpending     string     `json:"min_time_spending"`
	PlUsd               float64    `json:"pl_usd"`
	ReceivingActivity   []Activity `json:"receiving_activity"`
	SpendingActivity    []Activity `json:"spending_activity"`
}

type Activity struct {
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Value string `json:"value"`
}

// GetAddress
// https://api.blockchair.com/bitcoin/dashboards/address/{address}
func (c *Client) GetAddress(address string) (a *Address, e error) {
	response := &AddressResponse{}
	e = c.DoRequest("/dashboards/address/"+address, response)

	if len(response.Data) == 1 {
		a = &response.Data[0]
	} else {
		if len(response.Data) > 1 {
			e = fmt.Errorf("Unexpected response from the server")
		}
	}

	return
}
