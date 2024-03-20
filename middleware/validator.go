package middleware

import (
	"net/http"

	"github.com/fazriegi/my-gram/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type inputData interface {
	model.User | model.UserLoginRequest | model.UserUpdateRequest
}

type ValidationErrResponse struct {
	FailedField string
	Tag         string
	TagValue    string
}

func Validate[T inputData](data T) []*ValidationErrResponse {
	validate := validator.New()
	err := validate.Struct(data)
	var validationErrors []*ValidationErrResponse

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var validationErr ValidationErrResponse
			validationErr.FailedField = err.Field()
			validationErr.Tag = err.Tag()
			validationErr.TagValue = err.Param()
			validationErrors = append(validationErrors, &validationErr)
		}
	}

	return validationErrors
}

func ValidateField[T inputData]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := new(T)

		if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "error parsing request body",
			})
			return
		}

		if validationErrors := Validate(*data); validationErrors != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"message": "validation error",
				"errors":  validationErrors,
			})
			return
		}

		ctx.Next()
	}
}
