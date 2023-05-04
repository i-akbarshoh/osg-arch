package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/i-akbarshoh/osg-arch/internal/middleware"
)

func New() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.Authorizer())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}
