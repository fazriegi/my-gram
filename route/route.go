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
	c.SetupPhotoRoutes()
	c.SetupCommentRoutes()
	c.SetupSocialMediaRoutes()
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

func (c *RouteConfig) SetupPhotoRoutes() {
	photoRepository := repository.NewPhotoRepository(c.DB)
	photoUsecase := usecase.NewPhotoUsecase(photoRepository, c.Logger)
	photoController := controller.NewPhotoController(photoUsecase, c.Logger)

	photos := c.App.Group("/photos")
	photos.Use(middleware.Authentication())
	photos.POST("/", middleware.ValidateField[model.PhotoRequest](), photoController.Create)
	photos.GET("/", photoController.GetAll)
	photos.PUT("/:photoId", middleware.PhotoAuthorization(),
		middleware.ValidateField[model.PhotoRequest](), photoController.Update)
	photos.DELETE("/:photoId", middleware.PhotoAuthorization(), photoController.Delete)
}

func (c *RouteConfig) SetupCommentRoutes() {
	commentRepository := repository.NewCommentRepository(c.DB)
	commentUsecase := usecase.NewCommentUsecase(commentRepository, c.Logger)
	commentController := controller.NewCommentController(commentUsecase, c.Logger)

	comments := c.App.Group("/comments")
	comments.Use(middleware.Authentication())
	comments.POST("/", middleware.ValidateField[model.CreateCommentRequest](), commentController.Create)
	comments.GET("/", commentController.GetAllByUserId)
	comments.PUT("/:commentId", middleware.ValidateField[model.UpdateCommentRequest](),
		middleware.CommentAuthorization(), commentController.Update)
	comments.DELETE("/:commentId", middleware.CommentAuthorization(), commentController.Delete)
}

func (c *RouteConfig) SetupSocialMediaRoutes() {
	socialMediaRepository := repository.NewSocialMediaRepository(c.DB)
	socialMediaUsecase := usecase.NewSocialMediaUsecase(socialMediaRepository, c.Logger)
	socialMediaController := controller.NewSocialMediaController(socialMediaUsecase, c.Logger)

	socialMedias := c.App.Group("/socialmedias")
	socialMedias.Use(middleware.Authentication())
	socialMedias.POST("/", middleware.ValidateField[model.SocialMediaRequest](),
		socialMediaController.Create)
	socialMedias.GET("/", socialMediaController.GetAll)
	socialMedias.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(),
		middleware.ValidateField[model.SocialMediaRequest](), socialMediaController.Update)
	socialMedias.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(),
		socialMediaController.Delete)
}
