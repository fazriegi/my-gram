package repository

import (
	"github.com/fazriegi/my-gram/model"
	"gorm.io/gorm"
)

type IPhotoRepository interface {
	Create(props *model.Photo) error
	GetAll() ([]model.Photo, error)
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

func (r *PhotoRepository) GetAll() ([]model.Photo, error) {
	var photo []model.Photo

	err := r.db.
		Preload("User").
		Find(&photo).
		Error

	if err != nil {
		return nil, err
	}

	return photo, nil
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
