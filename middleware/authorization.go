package middleware

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/fazriegi/my-gram/config"
	"github.com/fazriegi/my-gram/model"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := config.GetDB()

		photoId, _ := strconv.Atoi(ctx.Param("photoId"))
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(float64))
		photo := model.Photo{}

		err := db.Select("user_id").First(&photo, photoId).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "data not found",
			})
			return
		}

		if photo.UserId != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		ctx.Next()
	}
}
