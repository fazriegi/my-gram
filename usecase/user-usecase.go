package usecase

import (
	"errors"

	"github.com/fazriegi/my-gram/libs/helper"
	"github.com/fazriegi/my-gram/model"
	"github.com/fazriegi/my-gram/repository"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	repository repository.IUserRepository
	log        *logrus.Logger
}

func NewUserUsecase(userRepository repository.IUserRepository, logger *logrus.Logger) *UserUsecase {
	return &UserUsecase{repository: userRepository, log: logger}
}

func (u *UserUsecase) Create(props *model.User) (*model.UserResponse, error) {
	var err error

	_, err = u.repository.FindByEmail(props.Email)

	if err != nil && err != gorm.ErrRecordNotFound {
		u.log.Errorf("error finding email: %s", err.Error())
		return nil, errors.New("failed to create user")
	} else if err == nil {
		return nil, errors.New("email already exist")
	}

	_, err = u.repository.FindByUsername(props.Username)

	if err != nil && err != gorm.ErrRecordNotFound {
		u.log.Errorf("error finding username: %s", err.Error())
		return nil, errors.New("failed to create user")
	} else if err == nil {
		return nil, errors.New("username already exist")
	}

	if props.Password, err = helper.HashPassword(props.Password); err != nil {
		u.log.Errorf("error hashing password: %s", err.Error())
		return nil, errors.New("failed to create user")
	}

	if err := u.repository.Create(props); err != nil {
		u.log.Errorf("error creating user: %s", err.Error())
		return nil, errors.New("failed to create user")
	}

	data := model.UserResponse{
		ID:       props.ID,
		Email:    props.Email,
		Username: props.Username,
		Age:      props.Age,
	}

	return &data, nil
}

func (u *UserUsecase) Login(props *model.UserLoginRequest) (string, error) {
	user, err := u.repository.FindByEmail(props.Email)

	if err != nil && err == gorm.ErrRecordNotFound {
		return "", errors.New("unauthorized")
	} else if err != nil {
		u.log.Errorf("error finding email")
		return "", err
	}

	isAuthenticate := helper.CheckPasswordHash(props.Password, user.Password)

	if !isAuthenticate {
		return "", errors.New("unauthorized")
	}

	token, err := helper.GenerateToken(user.ID, user.Email)

	if err != nil {
		u.log.Errorf("error generating token: %s", err.Error())
		return "", errors.New("unexpected error occured")
	}

	return token, nil
}

func (u *UserUsecase) Update(id int, props model.UserUpdateRequest) (*model.UserUpdateResponse, error) {
	existingUser, err := u.repository.FindByEmail(props.Email)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingUser.ID != 0 && existingUser.ID != id {
		return nil, errors.New("email already used")
	}

	existingUser, err = u.repository.FindByUsername(props.Username)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingUser.ID != 0 && existingUser.ID != id {
		return nil, errors.New("username already used")
	}

	user, err := u.repository.Update(id, &props)

	if err != nil {
		u.log.Errorf("error updating user: %s", err.Error())
		return nil, errors.New("unexpected error occured")
	}

	userResponse := model.UserUpdateResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}

	return &userResponse, nil
}

func (u *UserUsecase) Delete(id int) error {
	err := u.repository.Delete(id)

	if err != nil {
		u.log.Errorf("error deleting user account: %s", err.Error())
		return errors.New("unexpected error occured")
	}

	return nil
}
