package api

import "errors"

var (
	// ErrInvalidCredentials is given when the requestor provides invalid user login credentials
	ErrInvalidCredentials = errors.New("Invalid username and/or password")
)
