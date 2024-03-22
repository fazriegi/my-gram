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

type SocialMediaController struct {
	Log     *logrus.Logger
	UseCase *usecase.SocialMediaUsecase
}

func NewSocialMediaController(useCase *usecase.SocialMediaUsecase, logger *logrus.Logger) *SocialMediaController {
	return &SocialMediaController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *SocialMediaController) Create(ctx *gin.Context) {
	var socialmedia model.SocialMedia
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialmedia.UserId = int(userData["id"].(float64))

	if err := ctx.ShouldBindBodyWith(&socialmedia, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing request body",
		})
		return
	}

	created, err := c.UseCase.Create(&socialmedia)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success add social media",
		"data":    created,
	})
}

func (c *SocialMediaController) GetAll(ctx *gin.Context) {
	data, err := c.UseCase.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success fetch social media",
		"data":    data,
	})
}

func (c *SocialMediaController) Update(ctx *gin.Context) {
	var socialmedia model.SocialMedia
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialmedia.ID, _ = strconv.Atoi(ctx.Param("socialMediaId"))
	socialmedia.UserId = int(userData["id"].(float64))

	if err := ctx.ShouldBindBodyWith(&socialmedia, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing request body",
		})
		return
	}

	data, err := c.UseCase.Update(&socialmedia)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success update social media",
		"data":    data,
	})
}

func (c *SocialMediaController) Delete(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))

	err := c.UseCase.Delete(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}
