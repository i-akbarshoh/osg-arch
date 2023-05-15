package project

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/i-akbarshoh/osg-arch/internal/entity"
	"github.com/i-akbarshoh/osg-arch/internal/usecase/project"
)
	

type Controller struct {
	useCase *project.UseCase
}

func New(u *project.UseCase) *Controller {
	return &Controller{useCase: u}
}

func (con *Controller) CreateTask(c *gin.Context) {
	var body entity.Task
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind json, " + err.Error(),
		})
		return
	}

	ctx := context.TODO()
	id, err := con.useCase.CreateTask(ctx, body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot create task, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

func (con *Controller) Create(c *gin.Context) {
	var body entity.Project
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind json, " + err.Error(),
		})
		return
	}

	ctx := context.TODO()
	id, err := con.useCase.Create(ctx, body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot create project, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

func (con *Controller) List(c *gin.Context) {
	ctx := context.TODO()
	list, err := con.useCase.List(ctx)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot list project, " + err.Error(),
		})
		return
	}

	c.JSON(200, list)
}

func (con *Controller) Update(c *gin.Context) {
	var body entity.Project
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind json, " + err.Error(),
		})
		return
	}

	ctx := context.TODO()
	if err := con.useCase.Update(ctx, body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot update project, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":"success",
	})
}

func (con *Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot parse id to int, " + err.Error(),
		})
		return
	}

	ctx := context.TODO()
	if err := con.useCase.Delete(ctx, ID); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot delete project, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":"success",
	})
}