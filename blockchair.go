// Copyright 2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package blockchair Golang client Blockchair api.
//
// https://github.com/Blockchair/Blockchair.Support/blob/master/API.md
package blockchair // import "github.com/xorcare/blockchair"

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// Version api client.
	Version = "0.4"

	//Package package name const
	Package = "blockchair"

	// UserAgent is the header string used to identify this package.
	UserAgent = Package + "-api-go-client/" + Version + " (go; github; +https://git.io/fAJhJ)"

	// BasePath the root address in the network.
	BasePath = "https://api-v1.blockchair.com"
)

// Currency special type for currency codes like enum.
type Currency uint8

// String returns a string representation for the currency.
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
	// Bitcoin constant storing bitcoin cryptocurrency code for use inside the package.
	Bitcoin = Currency(0)
	// BitcoinCash constant storing bitcoin cash cryptocurrency code for use inside the package.
	BitcoinCash = Currency(1)
)

// Client specifies the mechanism by which individual API requests are made.
type Client struct {
	client    *http.Client
	BasePath  string   // API endpoint base URL.
	UserAgent string   // Optional additional User-Agent fragment.
	Currency  Currency // The currency in which information is requested.
}

func (c *Client) userAgent() string {
	c.UserAgent = strings.TrimSpace(c.UserAgent)
	if c.UserAgent == "" {
		return UserAgent
	}

	return UserAgent + " " + c.UserAgent
}

// Do to send an client request, which is then converted to the passed type.
func (c *Client) Do(path string, i interface{}) error {
	req, e := http.NewRequest("GET", c.BasePath+"/"+c.Currency.String()+path, nil)
	if e != nil {
		return c.err2(ErrCGD, e)
	}
	req.Header.Set("User-Agent", c.userAgent())

	resp, e := c.client.Do(req)
	if e != nil {
		return c.err3(ErrCGD, e, resp)
	}
	defer resp.Body.Close()

	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return c.err3(ErrCRR, e, resp)
	}

	if resp.Status[0] != '2' {
		return c.err3(ErrIRS, errors.New(string(bytes)), resp)
	}

	if e = json.Unmarshal(bytes, &i); e != nil {
		return c.err3(ErrRPE, e, resp)
	}

	return nil
}

// New specifies the mechanism by create new client the network internet.
func New(u Currency) *Client {
	return &Client{client: &http.Client{}, BasePath: BasePath, Currency: u}
}

// SetClient client client setter.
func (c *Client) SetClient(client *http.Client) {
	if client == nil {
		panic(Package + ": impossible install the client as nil")
	}
	c.client = client
}
