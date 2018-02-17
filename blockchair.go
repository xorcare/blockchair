// Copyright 2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Golang client Blockchair api -> https://github.com/Blockchair/Blockchair.Support/blob/master/API.md

package blockchair

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// APIRootNet the root address in the network
	APIRootNet = "https://api.blockchair.com"
)

const (
	Bitcoin = iota
	BitcoinCash
)

var currencyURL = map[uint8]string{
	Bitcoin:     "bitcoin",
	BitcoinCash: "bitcoin-cash",
}

// Client specifies the mechanism by which individual API requests are made.
type Client struct {
	http     *http.Client
	apiRoot  string
	currency uint8
}

// DoRequest to send an http request, which is then converted to the passed type.
func (c *Client) DoRequest(path string, i interface{}) (e error) {
	currencyURLString := currencyURL[c.currency]
	fullPath := c.apiRoot + "/" + currencyURLString + path
	response, e := c.http.Get(fullPath)
	if e != nil {
		return
	}

	defer response.Body.Close()

	bytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return
	}
	if response.Status[0] != '2' {
		return fmt.Errorf("Response error status %3s: %s", response.Status, string(bytes))
	}

	return json.Unmarshal(bytes, &i)
}

// New specifies the mechanism by create new client the network internet
func New(u uint8) *Client {
	return &Client{http: &http.Client{}, apiRoot: APIRootNet, currency: u}
}

// SetCurrency Currency setter
func (c *Client) SetCurrency(u uint8) {
	c.currency = u
}

// SetHTTP http client setter
func (c *Client) SetHTTP(cli *http.Client) {
	c.http = cli
}
