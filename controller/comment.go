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

type CommentController struct {
	Log     *logrus.Logger
	UseCase *usecase.CommentUsecase
}

func NewCommentController(useCase *usecase.CommentUsecase, logger *logrus.Logger) *CommentController {
	return &CommentController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *CommentController) Create(ctx *gin.Context) {
	var comment model.Comment
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	comment.UserId = int(userData["id"].(float64))

	if err := ctx.ShouldBindBodyWith(&comment, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing request body",
		})
		return
	}

	created, err := c.UseCase.Create(&comment)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success add comment",
		"data":    created,
	})
}

func (c *CommentController) GetAll(ctx *gin.Context) {
	data, err := c.UseCase.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success fetch comments",
		"data":    data,
	})
}

func (c *CommentController) Update(ctx *gin.Context) {
	var reqBody model.UpdateCommentRequest
	reqBody.ID, _ = strconv.Atoi(ctx.Param("commentId"))

	if err := ctx.ShouldBindBodyWith(&reqBody, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing request body",
		})
		return
	}

	comment := model.Comment{
		ID:      reqBody.ID,
		Message: reqBody.Message,
	}

	data, err := c.UseCase.Update(&comment)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success update comment",
		"data":    data,
	})
}

func (c *CommentController) Delete(ctx *gin.Context) {
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
		"message": "your comment has been successfully deleted",
	})
}
