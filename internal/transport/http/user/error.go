package user

import "errors"

var (
	ErrValidation = errors.New("data validation error")
	ErrBind       = errors.New("bind error")
	ErrUserExists = errors.New("user alredy exists")
)
