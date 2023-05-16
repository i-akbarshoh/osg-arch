package user

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/i-akbarshoh/osg-arch/internal/entity"
	"github.com/i-akbarshoh/osg-arch/internal/pkg/jwt"
	"github.com/i-akbarshoh/osg-arch/internal/usecase/user"
)

type Controller struct {
	useCase *user.UseCase
}

func NewController(useCase *user.UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}

func (con *Controller) Register(c *gin.Context) {
	var (
		body entity.User
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind json, " + err.Error(),
		})
		return
	}

	id, err := con.useCase.Register(c, body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot register user, " + err.Error(),
		})
		return
	}
	tokens, err := jwt.GenerateNewTokens(id, map[string]string{"role": body.Position})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot generate tokens, " + err.Error(),
		})
		return
	}

	c.JSON(200, entity.RegisterResponse{
		ID:     id,
		Tokens: *tokens,
	})
}

func (con *Controller) Login(c *gin.Context) {
	var (
		body entity.User
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind json, " + err.Error(),
		})
		return
	}
	ctx := context.Background()
	if err := con.useCase.Login(ctx, entity.User{
		ID: body.ID,
		Password: body.Password,
	}); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot login, " + err.Error(),
		})
		return
	}

	tokens, err := jwt.GenerateNewTokens(body.ID, map[string]string{"role": body.Position})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot generate tokens, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"access": tokens.Access,
		"expire": tokens.AccExpire,
		"refresh": tokens.Refresh,
	})
}

func (con *Controller) List(c *gin.Context) {
	ctx := context.Background()
	res, err := con.useCase.List(ctx)
	if  err != nil {
		c.JSON(400, gin.H{
			"message": "cannot list users, " + err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (con *Controller) Get(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	res, err := con.useCase.Get(ctx, id)
	if  err != nil {
		c.JSON(400, gin.H{
			"message": "cannot get user, " + err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (con *Controller) Update(c *gin.Context) {
	var body entity.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind json, " + err.Error(),
		})
		return
	}

	ctx := context.Background()
	if err := con.useCase.Update(ctx, body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot update user, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"masseage": "successs",
	})
}

func (con *Controller) Delete(c *gin.Context) {
	var body entity.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind json, " + err.Error(),
		})
		return
	}
	ctx := context.Background()
	if err := con.useCase.Delete(ctx, body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot delete user, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"masseage": "successs",
	})
}

func (con *Controller) CreateAttendance(c *gin.Context) {
	var body entity.UserAttendance
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind json, " + err.Error(),
		})
		return
	}

	if err := con.useCase.CreateAttendance(context.Background(), body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot create attendance, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"masseage": "successs",
	})
}

func (con *Controller) ListAttendance(c *gin.Context) {
	user_id := c.Param("user_id")
	ty := c.Param("type")
	list, err := con.useCase.ListAttendance(context.Background(), user_id, ty)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot list attendance, " + err.Error(),
		})
		return
	}

	c.JSON(200, list)
}