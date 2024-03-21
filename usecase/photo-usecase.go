package usecase

import (
	"errors"

	"github.com/fazriegi/my-gram/model"
	"github.com/fazriegi/my-gram/repository"
	"github.com/sirupsen/logrus"
)

type PhotoUsecase struct {
	repository repository.IPhotoRepository
	log        *logrus.Logger
}

func NewPhotoUsecase(PhotoRepository repository.IPhotoRepository, logger *logrus.Logger) *PhotoUsecase {
	return &PhotoUsecase{repository: PhotoRepository, log: logger}
}

func (u *PhotoUsecase) Create(props *model.Photo) (*model.CreatePhotoResponse, error) {
	if err := u.repository.Create(props); err != nil {
		u.log.Errorf("error adding photo: %s", err.Error())
		return nil, errors.New("unexpected error occured")
	}

	data := model.CreatePhotoResponse{
		ID:        props.ID,
		Title:     props.Title,
		Caption:   props.Caption,
		PhotoUrl:  props.PhotoUrl,
		UserId:    props.UserId,
		CreatedAt: props.CreatedAt,
	}

	return &data, nil
}

func (u *PhotoUsecase) GetAllByUserId(userId int) ([]model.GetAllPhotoByUserResponse, error) {
	data, err := u.repository.GetAllByUserId(userId)

	if err != nil {
		u.log.Errorf("error fetching photo: %s", err.Error())
		return nil, errors.New("unexpected error occured")
	}

	var response []model.GetAllPhotoByUserResponse

	for _, v := range *data {
		photoResponse := model.GetAllPhotoByUserResponse{
			ID:        v.ID,
			Title:     v.Title,
			Caption:   v.Caption,
			PhotoUrl:  v.PhotoUrl,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: struct {
				Email    string "json:\"email\""
				Username string "json:\"username\""
			}{
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		}

		response = append(response, photoResponse)
	}

	return response, nil
}

func (u *PhotoUsecase) Update(props *model.Photo) (*model.UpdatePhotoResponse, error) {
	err := u.repository.Update(props)

	if err != nil {
		u.log.Errorf("error updating user: %s", err.Error())
		return nil, errors.New("unexpected error occured")
	}

	response := model.UpdatePhotoResponse{
		ID:        props.ID,
		Title:     props.Title,
		Caption:   props.Caption,
		PhotoUrl:  props.PhotoUrl,
		UserId:    props.UserId,
		UpdatedAt: props.UpdatedAt,
	}

	return &response, nil
}

func (u *PhotoUsecase) Delete(id int) error {
	err := u.repository.Delete(id)

	if err != nil {
		u.log.Errorf("error deleting photo: %s", err.Error())
		return errors.New("unexpected error occured")
	}

	return nil
}
