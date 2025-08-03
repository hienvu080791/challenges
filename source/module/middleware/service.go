package middleware

import (
	"demo_challenges/source/module/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Service interface {
	Middleware(mustAuth bool, api func(*gin.Context)) gin.HandlerFunc
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Middleware(mustAuth bool, api func(*gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if mustAuth {
			accessToken := c.GetHeader("Authorization")
			if len(accessToken) == 0 || !strings.HasPrefix(accessToken, "Bearer ") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				return
			}
			//validate the access token
			valid := utils.ValidateToken(accessToken[7:])
			if !valid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid access token"})
				return
			}
		}
		api(c)
	}
}
