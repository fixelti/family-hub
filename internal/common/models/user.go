package models

import "time"

type UserDTO struct {
	ID       uint   `json:"id"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID        uint `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"updated_at"`
	DeleteAt  *time.Time `db:"deleted_at"`
}

func (user User) ToUserDTO() UserDTO {
	return UserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}
