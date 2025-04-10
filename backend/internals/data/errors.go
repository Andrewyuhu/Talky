package data

import "errors"

var (
	ErrDuplicateUsername  = errors.New("USERNAME_EXISTS")
	ErrRecordNotFound     = errors.New("RECORD_NOT_FOUND")
	ErrDuplicateEmail     = errors.New("EMAIL_EXISTS")
	ErrInvalidCredentials = errors.New("INVALID_CREDENTIALS")
	ErrDuplicateChat      = errors.New("CHAT_EXISTS")
)
