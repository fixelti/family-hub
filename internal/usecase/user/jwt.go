package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)
 
func (user userUsecase) SignUp(ctx context.Context, email, password string) (uint, error) {
	userID, err := user.db.GetByEmail(ctx, email)
	if err != nil {
		return 0, err
	}
	if userID != 0 {
		return 0, ErrUserExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return 0, err
	}

	id, err := user.db.Create(ctx, email, string(passwordHash))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func tokenValidation(tokenKey string) (bool, error) { return false, nil }

func createToken(tokenKey string, tokenLifeTime uint, paylaod map[string]any) (string, error) {
	return "", nil
}
