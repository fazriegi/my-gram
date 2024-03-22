package repository

import (
	"github.com/fazriegi/my-gram/model"
	"gorm.io/gorm"
)

type ISocialMediaRepository interface {
	Create(props *model.SocialMedia) error
	GetAll() ([]model.SocialMedia, error)
	Update(props *model.SocialMedia) error
	Delete(id int) error
}

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{db}
}

func (r *SocialMediaRepository) Create(props *model.SocialMedia) error {
	return r.db.Create(props).Error
}

func (r *SocialMediaRepository) GetAll() ([]model.SocialMedia, error) {
	var socialmedia []model.SocialMedia

	err := r.db.
		Preload("User").
		Find(&socialmedia).
		Error

	if err != nil {
		return nil, err
	}

	return socialmedia, nil
}

func (r *SocialMediaRepository) Update(props *model.SocialMedia) error {
	tx := r.db.Begin()
	err := tx.Where("id = ?", props.ID).Updates(&props).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *SocialMediaRepository) Delete(id int) error {
	tx := r.db.Begin()

	if err := tx.Delete(&model.SocialMedia{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
