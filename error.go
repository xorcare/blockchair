package blockchair

import (
	"errors"
	"net/http"
)

// Errors it is a set of errors returned when working with the package.
var (
	ErrAIW = errors.New("address is wrong")
	ErrCGD = errors.New("cannot get data on url")
	ErrCRR = errors.New("could not read answer response")
	ErrIRS = errors.New("incorrect response status")
	ErrRPE = errors.New("response parsing error")
)

// Error data structure describing the error.
type Error struct {
	// Address wrong address.
	Address *string
	// ErrMain error information from the standard package error set.
	ErrMain error
	// ErrExec information about the error that occurred during
	// the operation of the standard library or external packages.
	ErrExec error
	// Response client response.
	Response *http.Response
}

func (c *Client) error(errorMain error) error {
	return c.newError(errorMain, nil, nil, nil)
}

func (c *Client) errors(errorMain error, errorExec error) error {
	return c.newError(errorMain, errorExec, nil, nil)
}

func (c *Client) errorAddress(errorMain error, address string) error {
	return c.newError(errorMain, nil, nil, &address)
}

func (c *Client) errorResponse(errorMain error, errorExec error, response *http.Response) error {
	return c.newError(errorMain, errorExec, response, nil)
}

func (c *Client) newError(errorMain error, errorExec error, response *http.Response, address *string) error {
	if errorMain == nil {
		return nil
	}

	return &Error{
		ErrMain:  errorMain,
		ErrExec:  errorExec,
		Response: response,
		Address:  address,
	}
}

// Error compatibility with error interface.
func (e Error) Error() string {
	return e.ErrMain.Error()
}
