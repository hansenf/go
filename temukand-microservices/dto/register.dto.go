package dto

type RegisterRequest struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required" validate:"email"`
	Password      string `json:"password"  binding:"required" validate:"min:8"`
	CheckPassword string `json:"checkPassword"  binding:"required" validate:"min:8"`
}
