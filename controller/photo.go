package controller

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/fazriegi/my-gram/model"
	"github.com/fazriegi/my-gram/usecase"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

type PhotoController struct {
	Log     *logrus.Logger
	UseCase *usecase.PhotoUsecase
}

func NewPhotoController(useCase *usecase.PhotoUsecase, logger *logrus.Logger) *PhotoController {
	return &PhotoController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *PhotoController) Create(ctx *gin.Context) {
	var photo model.Photo
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	photo.UserId = int(userData["id"].(float64))

	if err := ctx.ShouldBindBodyWith(&photo, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing request body",
		})
		return
	}

	created, err := c.UseCase.Create(&photo)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success add photo",
		"data":    created,
	})
}

func (c *PhotoController) GetAll(ctx *gin.Context) {
	data, err := c.UseCase.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success fetch photos",
		"data":    data,
	})
}

func (c *PhotoController) Update(ctx *gin.Context) {
	var photo model.Photo
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	photo.ID, _ = strconv.Atoi(ctx.Param("photoId"))
	photo.UserId = int(userData["id"].(float64))

	if err := ctx.ShouldBindBodyWith(&photo, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing request body",
		})
		return
	}

	data, err := c.UseCase.Update(&photo)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success update photo",
		"data":    data,
	})
}

func (c *PhotoController) Delete(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	err := c.UseCase.Delete(photoId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your photo has been successfully deleted",
	})
}
