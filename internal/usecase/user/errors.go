package user

import "errors"

var (
	ErrUserExists = errors.New("user alredy exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
