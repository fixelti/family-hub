package models

import "time"

type UserDTO struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        uint
	Email     string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
}

func (user User) ToUserDTO() UserDTO {
	return UserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}
