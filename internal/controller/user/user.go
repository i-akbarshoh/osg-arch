package user

import (
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
