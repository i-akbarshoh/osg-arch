package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/i-akbarshoh/osg-arch/internal/controller"
	"github.com/i-akbarshoh/osg-arch/internal/middleware"
)

func New(c controller.New) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.Authorizer())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	r.POST("/register", c.Register)
	return r
}
