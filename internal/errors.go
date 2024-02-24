package customError

import "errors"

// repository errors
var (
	ErrBeginTransaction = errors.New("DB-000001")
	ErrMakeQuery        = errors.New("DB-000002")
	ErrScanInSctruct    = errors.New("DB-000002")
)

// usecase errors
var (
	ErrInternal           = errors.New("US-000001")
	ErrUserExists         = errors.New("US-000002")
	ErrInvalidCredentials = errors.New("US-000003")
	ErrTokenIsNotValid    = errors.New("US-000004")
	ErrBind               = errors.New("US-000005")
)
