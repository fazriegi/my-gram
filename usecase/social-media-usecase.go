package usecase

import (
	"errors"

	"github.com/fazriegi/my-gram/model"
	"github.com/fazriegi/my-gram/repository"
	"github.com/sirupsen/logrus"
)

type SocialMediaUsecase struct {
	repository repository.ISocialMediaRepository
	log        *logrus.Logger
}

func NewSocialMediaUsecase(SocialMediaRepository repository.ISocialMediaRepository, logger *logrus.Logger) *SocialMediaUsecase {
	return &SocialMediaUsecase{repository: SocialMediaRepository, log: logger}
}

func (u *SocialMediaUsecase) Create(props *model.SocialMedia) (*model.CreateSocialMediaResponse, error) {
	if err := u.repository.Create(props); err != nil {
		u.log.Errorf("error adding social media: %s", err.Error())
		return nil, errors.New("unexpected error occured")
	}

	data := model.CreateSocialMediaResponse{
		ID:             props.ID,
		Name:           props.Name,
		SocialMediaUrl: props.SocialMediaUrl,
		UserId:         props.UserId,
		CreatedAt:      props.CreatedAt,
	}

	return &data, nil
}

func (u *SocialMediaUsecase) GetAll() ([]model.GetAllSocialMediaResponse, error) {

	data, err := u.repository.GetAll()

	if err != nil {
		u.log.Errorf("error fetching social media: %s", err.Error())
		return nil, errors.New("unexpected error occured")
	}

	var response []model.GetAllSocialMediaResponse

	for _, v := range data {
		socialmediaResponse := model.GetAllSocialMediaResponse{
			ID:             v.ID,
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			UserId:         v.UserId,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			User: struct {
				ID       int    "json:\"id\""
				Email    string "json:\"email\""
				Username string "json:\"username\""
			}{
				ID:       v.ID,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		}

		response = append(response, socialmediaResponse)
	}

	return response, nil

}

func (u *SocialMediaUsecase) Update(props *model.SocialMedia) (model.UpdateSocialMediaResponse, error) {
	err := u.repository.Update(props)

	if err != nil {
		u.log.Errorf("error updating social media: %s", err.Error())
		return model.UpdateSocialMediaResponse{}, errors.New("unexpected error occured")
	}

	response := model.UpdateSocialMediaResponse{
		ID:             props.ID,
		Name:           props.Name,
		SocialMediaUrl: props.SocialMediaUrl,
		UserId:         props.UserId,
		UpdatedAt:      props.UpdatedAt,
	}

	return response, nil
}

func (u *SocialMediaUsecase) Delete(id int) error {
	err := u.repository.Delete(id)

	if err != nil {
		u.log.Errorf("error deleting social media: %s", err.Error())
		return errors.New("unexpected error occured")
	}

	return nil
}
