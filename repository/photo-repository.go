package repository

import (
	"github.com/fazriegi/my-gram/model"
	"gorm.io/gorm"
)

type IPhotoRepository interface {
	Create(props *model.Photo) error
	GetAllByUserId(userId int) (*[]model.Photo, error)
	Update(props *model.Photo) error
	Delete(id int) error
}

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{db}
}

func (r *PhotoRepository) Create(props *model.Photo) error {
	return r.db.Create(props).Error
}

func (r *PhotoRepository) GetAllByUserId(userId int) (*[]model.Photo, error) {
	var photo []model.Photo

	err := r.db.Table("photos").
		Joins("JOIN users ON users.id = photos.user_id").
		Select("photos.*, users.username, users.email").
		Where("photos.user_id = ?", userId).
		Preload("User").
		Find(&photo)

	if err.Error != nil {
		return nil, err.Error
	}

	return &photo, nil
}

func (r *PhotoRepository) Update(props *model.Photo) error {
	tx := r.db.Begin()
	err := tx.Where("id = ?", props.ID).Updates(&props).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *PhotoRepository) Delete(id int) error {
	tx := r.db.Begin()

	if err := tx.Delete(&model.Photo{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
