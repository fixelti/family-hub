package user

import (
	"context"

	"github.com/fixelti/family-hub/internal/common/models"
	"golang.org/x/crypto/bcrypt"
)

func (user userUsecase) SignUp(ctx context.Context, email, password string) (models.UserDTO, error) {
	emailExists, err := user.db.GetByEmail(ctx, email)
	if err != nil {
		return models.UserDTO{}, err
	}
	if emailExists.ID != 0 {
		return models.UserDTO{}, ErrUserExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return models.UserDTO{}, err
	}

	createdUser, err := user.db.Create(ctx, email, string(passwordHash))
	if err != nil {
		return models.UserDTO{}, err
	}

	return createdUser.ToUserDTO(), nil
}

func tokenValidation(tokenKey string) (bool, error) {return false, nil}

func createToken(tokenKey string, tokenLifeTime uint, paylaod map[string]any) (string, error) {return "", nil}