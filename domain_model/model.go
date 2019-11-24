package model

import "errors"

var (
	ErrAccountNotFound       = errors.New("err account not found")
	ErrDuplicateEventVersion = errors.New("event version already exists")
	ErrStreamNotFound        = errors.New("stream not found")
)
