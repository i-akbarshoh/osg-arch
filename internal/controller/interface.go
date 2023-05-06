package controller

import "github.com/gin-gonic/gin"

type New interface {
	User
}

type User interface {
	Register(c *gin.Context)
}
