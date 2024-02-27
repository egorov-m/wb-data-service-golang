package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	auth "wb-data-service-golang/wb-data-service/internal/module/authorization/core"
)

func NewAuthMiddleware(authModule auth.AuthorizationUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Authorization token is required")
			return
		}
		splitted := strings.Split(auth, " ")
		userId, err := authModule.CheckToken(c, splitted[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		}

		c.Set("userId", userId)

		c.Next()
	}
}
