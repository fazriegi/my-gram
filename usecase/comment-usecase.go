package usecase

import (
	"errors"

	"github.com/fazriegi/my-gram/model"
	"github.com/fazriegi/my-gram/repository"
	"github.com/sirupsen/logrus"
)

type CommentUsecase struct {
	repository repository.ICommentRepository
	log        *logrus.Logger
}

func NewCommentUsecase(CommentRepository repository.ICommentRepository, logger *logrus.Logger) *CommentUsecase {
	return &CommentUsecase{repository: CommentRepository, log: logger}
}

func (u *CommentUsecase) Create(props *model.Comment) (*model.CreateCommentResponse, error) {
	if err := u.repository.Create(props); err != nil {
		u.log.Errorf("error adding comment: %s", err.Error())
		return nil, errors.New("unexpected error occured")
	}

	data := model.CreateCommentResponse{
		ID:        props.ID,
		Message:   props.Message,
		PhotoId:   props.PhotoId,
		UserId:    props.UserId,
		CreatedAt: props.CreatedAt,
	}

	return &data, nil
}

func (u *CommentUsecase) GetAll() ([]model.GetAllCommentsResponse, error) {
	data, err := u.repository.GetAll()

	if err != nil {
		u.log.Errorf("error fetching comments: %s", err.Error())
		return nil, errors.New("unexpected error occured")
	}

	var response []model.GetAllCommentsResponse

	for _, v := range data {
		commentResponse := model.GetAllCommentsResponse{
			ID:        v.ID,
			Message:   v.Message,
			PhotoId:   v.PhotoId,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: struct {
				ID       int    "json:\"id\""
				Email    string "json:\"email\""
				Username string "json:\"username\""
			}{
				ID:       v.User.ID,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
			Photo: struct {
				ID       int    "json:\"id\""
				Title    string "json:\"title\""
				Caption  string "json:\"caption\""
				PhotoUrl string "json:\"photo_url\""
				UserId   int    "json:\"user_id\""
			}{
				ID:       v.Photo.ID,
				Title:    v.Photo.Title,
				Caption:  v.Photo.Caption,
				PhotoUrl: v.Photo.PhotoUrl,
				UserId:   v.Photo.UserId,
			},
		}

		response = append(response, commentResponse)
	}

	return response, nil
}

func (u *CommentUsecase) Update(props *model.Comment) (model.UpdateCommentResponse, error) {
	err := u.repository.Update(props)

	if err != nil {
		u.log.Errorf("error updating comment: %s", err.Error())
		return model.UpdateCommentResponse{}, errors.New("unexpected error occured")
	}

	response := model.UpdateCommentResponse{
		ID:        props.ID,
		Message:   props.Message,
		UserId:    props.UserId,
		PhotoId:   props.PhotoId,
		UpdatedAt: props.UpdatedAt,
	}

	return response, nil
}

func (u *CommentUsecase) Delete(id int) error {
	err := u.repository.Delete(id)

	if err != nil {
		u.log.Errorf("error deleting comment: %s", err.Error())
		return errors.New("unexpected error occured")
	}

	return nil
}
