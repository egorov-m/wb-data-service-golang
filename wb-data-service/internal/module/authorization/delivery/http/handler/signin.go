package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"wb-data-service-golang/wb-data-service/internal/module/authorization/delivery/http/handler/request"
	token "wb-data-service-golang/wb-data-service/internal/module/token/core"
	user "wb-data-service-golang/wb-data-service/internal/module/user/core"
)

type AuthSignIn interface {
	SignIn(context.Context, user.User) (token.Token, error)
}

// @Summary Sign in
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param requestBody body request.SignInBody true "Sign in"
// @Success 200 {object} core.Token
// @Router /auth/sign-in [post]
func NewSignIn(useCase AuthSignIn) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody request.SignInBody

		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				gin.H{"message": "invalid request body"})

			return
		}

		token, err := useCase.SignIn(c, requestBody.ToEntity())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, token)
	}
}
