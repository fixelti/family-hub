package models

import "github.com/golang-jwt/jwt/v4"

type Tokens struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

type Payload struct {
	UserID   uint
	Expirate uint
	jwt.RegisteredClaims
}
