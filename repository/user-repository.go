package repository

import (
	"github.com/fazriegi/my-gram/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(props *model.User) error
	FindByEmail(email string) (model.User, error)
	FindByUsername(username string) (model.User, error)
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
