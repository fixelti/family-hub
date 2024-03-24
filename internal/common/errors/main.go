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
	ErrBind               = errors.New("US-000005")
)

// lib/jwt errors
var (
	ErrTokenWithClaimsIsNil = errors.New("JWT-000001")
	ErrTokenIsNotValid      = errors.New("JWT-000002")
)
