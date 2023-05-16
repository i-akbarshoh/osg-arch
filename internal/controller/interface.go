package controller

import "github.com/gin-gonic/gin"

type New interface {
	User
	Project
}

type User interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	CreateAttendance(c *gin.Context)
	ListAttendance(c *gin.Context)
}

type Project interface{
	CreateTask(c *gin.Context)
	Create(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	UpdateTask(c *gin.Context)
	GetTask(c *gin.Context)
	ListTasks(c *gin.Context)
	DeleteTask(c *gin.Context)
	CreateComment(c *gin.Context)
	ListComments(c *gin.Context)
	DeleteComment(c *gin.Context)
}
