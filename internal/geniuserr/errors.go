package geniuserr

import "errors"

var (
	ErrDryMismatch = errors.New("resulting generation mismatch with existing file")
)
