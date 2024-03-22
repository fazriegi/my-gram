package repository

import (
	"github.com/fazriegi/my-gram/model"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	Create(props *model.Comment) error
	GetAll() ([]model.Comment, error)
	Update(props *model.Comment) error
	Delete(id int) error
	BulkDeleteByUser(userId int) error
	BulkDeleteByPhoto(photoId int) error
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) Create(props *model.Comment) error {
	return r.db.Create(props).Error
}

func (r *CommentRepository) GetAll() ([]model.Comment, error) {
	var comment []model.Comment

	err := r.db.Table("comments").
		Preload("User").
		Preload("Photo").
		Find(&comment).
		Error

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) Update(props *model.Comment) error {
	tx := r.db.Begin()
	err := tx.Where("id = ?", props.ID).Updates(&props).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	r.db.First(&props, props.ID)

	return nil
}

func (r *CommentRepository) Delete(id int) error {
	tx := r.db.Begin()

	if err := tx.Delete(&model.Comment{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *CommentRepository) BulkDeleteByUser(userId int) error {
	if err := r.db.Where("user_id = ?", userId).Delete(&model.Comment{}).Error; err != nil {
		r.db.Rollback()
		return err
	}

	return nil
}

func (r *CommentRepository) BulkDeleteByPhoto(photoId int) error {
	if err := r.db.Where("photo_id = ?", photoId).Delete(&model.Comment{}).Error; err != nil {
		r.db.Rollback()
		return err
	}

	return nil
}
