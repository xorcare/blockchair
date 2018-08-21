// Copyright 2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchair

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// BasePath the root address in the network
	BasePath = "https://api.blockchair.com"
)

// Currency special type for Currency codes like enum
type Currency uint8

func (c Currency) String() (currency string) {
	switch c {
	case Bitcoin:
		currency = "bitcoin"
	case BitcoinCash:
		currency = "bitcoin-cash"
	}
	return currency
}

const (
	// Bitcoin constant storing bitcoin cryptocurrency code for use inside the package
	Bitcoin = Currency(0)
	// BitcoinCash constant storing bitcoin cash cryptocurrency code for use inside the package
	BitcoinCash = Currency(1)
)

// Client specifies the mechanism by which individual API requests are made.
type Client struct {
	http     *http.Client
	BasePath string
	Currency Currency
}

// DoRequest to send an http request, which is then converted to the passed type.
func (c *Client) DoRequest(path string, i interface{}) (e error) {
	fullPath := c.BasePath + "/" + c.Currency.String() + path
	response, e := c.http.Get(fullPath)
	if e != nil {
		return
	}

	defer response.Body.Close()
	bytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return
	}
	fmt.Println(string(bytes))
	if response.Status[0] != '2' {
		return fmt.Errorf("response error status %3s: %s", response.Status, string(bytes))
	}

	return json.Unmarshal(bytes, &i)
}

// New specifies the mechanism by create new client the network internet
func New(u Currency) *Client {
	return &Client{http: &http.Client{}, BasePath: BasePath, Currency: u}
}

// SetHTTP http client setter
func (c *Client) SetHTTP(client *http.Client) {
	c.http = client
}
