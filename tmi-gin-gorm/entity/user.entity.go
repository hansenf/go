package entity

import "time"

type User struct {
	ID        int64     `gorm:"primary_key:auto_increment" json:"-"`
	Name      string    `gorm:"type:varchar(100)" json:"-"`
	Email     string    `gorm:"type:varchar(100)" json:"-"`
	Password  string    `gorm:"type:varchar(100)" json:"-"`
	NomorHP   string    `gorm:"type:varchar(100)" json:"-"`
	Role      string    `gorm:"type:varchar(100)" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `gorm:"-" json:"token,omitempty"`
}
