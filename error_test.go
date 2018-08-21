// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchair

import (
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {
	if NewError(nil, nil, nil, nil) != nil {
		t.Fatal("Wrong error")
	}

	test := "test"
	resp := &http.Response{
		StatusCode: http.StatusOK,
	}
	err := NewError(ErrAIW, ErrCRR, nil, nil)
	if err.Response != nil {
		t.Fatal("wrong Error.Response expected nil: ", *err.Response)
	}
	if err.Address != nil {
		t.Fatal("wrong Error.Address expected nil: ", *err.Address)
	}

	err = NewError(ErrAIW, ErrCRR, resp, &test)
	if err.ErrMain != ErrAIW {
		t.Fatal("wrong Error.ErrMain: ", err.ErrMain)
	}
	if err.ErrExec != ErrCRR {
		t.Fatal("wrong Error.ErrExec", err.ErrExec)
	}
	if *err.Address != test {
		t.Fatal("wrong Error.Address", *err.Address)
	}
	if err.Response.StatusCode != http.StatusOK {
		t.Fatal("wrong Error.Response", *err.Response)
	}
}
