package middleware

import (
	"github.com/casbin/casbin/v2"
	_ "github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
	_ "gitlab.com/dotgo13/url-shortner/pkg/config"
	"log"
	"net/http"
)

func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		enforcer, err := casbin.NewEnforcer(config.Data()["config_path"], config.Data()["casbin_roles_path"])
		if err != nil {
			log.Fatal("enforcer not initialized, ", err)
			return
		}

		claims, err := extractClaims(accessToken, []byte(config.Data()["signing_key"].(string)))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		role := claims["role"]

		ok, err := enforcer.Enforce(role, c.Request.URL.String(), c.Request.Method)
		if err != nil {
			log.Println("could not enforce:", err)
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusForbidden, map[string]string{
				"error": "user not allowed, there is a problem with authorization",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
