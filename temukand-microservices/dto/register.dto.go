package dto

import "time"

type RegisterRequest struct {
	Name          string    `json:"name" binding:"required"`
	Email         string    `json:"email" binding:"required" validate:"email"`
	Password      string    `json:"password"  binding:"required" validate:"min:8"`
	CheckPassword string    `json:"checkPassword"  binding:"required" validate:"min:8"`
	NomorHP       string    `gorm:"type:varchar(16)" json:"nomor_hp"`
	Role          string    `gorm:"type:varchar(20)" json:"role"`
	CreatedAt     time.Time `json:"created_at"`
}
