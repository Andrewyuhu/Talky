package data

import "errors"

var (
	ErrDuplicateUsername  = errors.New("username already exists")
	ErrRecordNotFound     = errors.New("record not found")
	ErrDuplicateEmail     = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
