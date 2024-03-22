package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/fazriegi/my-gram/model"
	"github.com/fazriegi/my-gram/usecase"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUsecase
}

func NewUserController(useCase *usecase.UserUsecase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserController) Create(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing request body",
		})
		return
	}

	created, err := c.UseCase.Create(&user)

	if err != nil && strings.Contains(err.Error(), "exist") {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"data":    created,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	var user model.UserLoginRequest

	if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing request body",
		})
		return
	}

	token, err := c.UseCase.Login(&user)

	if err != nil && strings.Contains(err.Error(), "unauthorized") {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong email or password!",
		})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unexpected error occured",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success login",
		"token":   token,
	})
}

func (c *UserController) Update(ctx *gin.Context) {
	var user model.UserUpdateRequest
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	id, _ := strconv.Atoi(ctx.Param("id"))

	if userId != id {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	if err := ctx.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing request body",
		})
		return
	}

	data, err := c.UseCase.Update(userId, user)

	if err != nil && strings.Contains(err.Error(), "already used") {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success update user",
		"data":    data,
	})
}

func (c *UserController) Delete(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	id, _ := strconv.Atoi(ctx.Param("id"))

	if userId != id {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	err := c.UseCase.Delete(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your account has been successfully deleted",
	})
}
