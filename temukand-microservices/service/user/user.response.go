package _user

import (
	"time"
	"tmi-gin/entity"
)

type UserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token,omitempty"`
	NomorHP   string    `json:"nomor_hp"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     "",
		NomorHP:   user.NomorHP,
		Role:      "1",
		CreatedAt: user.CreatedAt,
	}
}
