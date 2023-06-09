package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/i-akbarshoh/osg-arch/internal/controller"
	"github.com/i-akbarshoh/osg-arch/internal/middleware"
)

func New(u controller.User, p controller.Project) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.Authorizer())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	r.POST("/register", u.Register)
	r.POST("/login", u.Login)
	r.GET("/list-user", u.List)
	r.GET("/get-user/:id", u.Get)
	r.PUT("/update-user", u.Update)
	r.DELETE("/delete-user", u.Delete)
	r.POST("/create-task", p.CreateTask)
	r.POST("/create-project", p.Create)
	r.GET("/list-project", p.List)
	r.PUT("/update-project", p.Update)
	r.DELETE("/delete-project/:id", p.Delete)
	r.PUT("/update-task", p.UpdateTask)
	r.GET("/get-task/:id", p.GetTask)
	r.GET("/list-tasks", p.ListTasks)
	r.DELETE("/delete-task/:id", p.DeleteTask)
	r.POST("/create-comment", p.CreateComment)
	r.GET("/list-comments/:id", p.ListComments)
	r.DELETE("/delete-comment/:id", p.DeleteComment)
	r.POST("/create-attendance", u.CreateAttendance)
	r.GET("/list-attendance/:user_id/:type", u.ListAttendance)
	return r
}
