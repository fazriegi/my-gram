package repository

import (
	"github.com/fazriegi/my-gram/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	Create(props *model.User) error
	FindByEmail(email string) (model.User, error)
	FindByUsername(username string) (model.User, error)
	Update(id int, props *model.UserUpdateRequest) (*model.User, error)
	Delete(id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(props *model.User) error {
	return r.db.Create(props).Error
}

func (r *UserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *UserRepository) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepository) Update(id int, props *model.UserUpdateRequest) (*model.User, error) {
	var user model.User
	tx := r.db.Begin()
	err := tx.Model(&user).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(props).
		Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &user, nil
}

func (r *UserRepository) Delete(id int) error {
	tx := r.db.Begin()

	if err := tx.Delete(&model.User{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
