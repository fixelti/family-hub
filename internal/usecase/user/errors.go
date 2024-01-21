package user

import "errors"

var (
	ErrUserExists = errors.New("user alredy exists")
)
