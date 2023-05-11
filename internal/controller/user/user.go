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