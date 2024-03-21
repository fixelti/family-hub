package crypto

import (
	libBcrypt "golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source ./main.go -destination ./mocks/main.go
type Crypto interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword []byte, password []byte) error
}

type crypto struct{}

func New() Crypto { return &crypto{} }

func (crypto) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return libBcrypt.GenerateFromPassword(password, cost)
}

func (crypto) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	return libBcrypt.CompareHashAndPassword(hashedPassword, password)
}
