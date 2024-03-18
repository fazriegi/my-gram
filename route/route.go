package route

import (
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

}
