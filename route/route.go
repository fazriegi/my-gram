package route

import (
	"github.com/fazriegi/my-gram/controller"
	"github.com/fazriegi/my-gram/middleware"
	"github.com/fazriegi/my-gram/model"
	"github.com/fazriegi/my-gram/repository"
	"github.com/fazriegi/my-gram/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RouteConfig struct {
	App    *gin.Engine
	DB     *gorm.DB
	Logger *logrus.Logger
}

func (c *RouteConfig) NewRoute() {
	c.SetupUserRoutes()
}

func (c *RouteConfig) SetupUserRoutes() {
	userRepository := repository.NewUserRepository(c.DB)
	userUsecase := usecase.NewUserUsecase(userRepository, c.Logger)
	userController := controller.NewUserController(userUsecase, c.Logger)

	users := c.App.Group("/users")

	users.POST("/register", middleware.ValidateField[model.User](), userController.Create)
	users.POST("/login", middleware.ValidateField[model.UserLoginRequest](), userController.Login)
	users.PUT("/:id", middleware.ValidateField[model.UserUpdateRequest](), middleware.Authentication(), userController.Update)
	users.DELETE("/:id", middleware.Authentication(), userController.Delete)
}
