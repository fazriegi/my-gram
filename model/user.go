package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"not null;index:idx_username" json:"username"`
	Email     string `gorm:"not null;index:idx_email" json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Age       uint8  `json:"age" validate:"gte=8"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint8  `json:"age"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
}

type UserUpdateResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       uint8     `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}
