package dto

type LoginRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:8"`
}
