package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	UserId    int    `json:"user_id"`
	PhotoId   int    `json:"photo_id"`
	Message   string `json:"message"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted"`
	User      User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserId;"`
	Photo     Photo          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:PhotoId;"`
}

type CreateCommentRequest struct {
	Message string `json:"message" validate:"required"`
	PhotoId int    `json:"photo_id"`
}

type CreateCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllCommentByUserResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"user"`
	Photo struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
		UserId   int    `json:"user_id"`
	} `json:"photo"`
}

type UpdateCommentRequest struct {
	ID      int    `json:"id"`
	Message string `json:"message" validate:"required"`
}

type UpdateCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
