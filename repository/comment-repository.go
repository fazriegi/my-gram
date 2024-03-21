package repository

import (
	"github.com/fazriegi/my-gram/model"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	Create(props *model.Comment) error
	GetAllByUserId(userId int) (*[]model.Comment, error)
	Update(props *model.Comment) error
	Delete(id int) error
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

func (r *CommentRepository) GetAllByUserId(userId int) (*[]model.Comment, error) {
	var comment []model.Comment

	err := r.db.Table("comments").
		Joins("JOIN users ON users.id = comments.user_id").
		Joins("JOIN photos ON photos.id = comments.photo_id").
		Where("comments.user_id = ?", userId).
		Preload("User").
		Preload("Photo").
		Find(&comment)

	if err.Error != nil {
		return nil, err.Error
	}

	return &comment, nil
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
