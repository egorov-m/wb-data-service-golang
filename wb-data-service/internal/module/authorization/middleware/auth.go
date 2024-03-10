package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wb-data-service-golang/wb-data-service/internal/domain"
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
		if len(splitted) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorInvalidToken)
			return
		}
		userId, err := authModule.CheckToken(c, splitted[1], "access")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}
