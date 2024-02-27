package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"wb-data-service-golang/wb-data-service/internal/module/authorization/delivery/http/handler/request"
	token "wb-data-service-golang/wb-data-service/internal/module/token/core"
	user "wb-data-service-golang/wb-data-service/internal/module/user/core"
)

type AuthSignUp interface {
	SignUp(context.Context, user.User) (token.Token, error)
}

func NewSignUp(useCase AuthSignUp) gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestBody request.SignUpBody

		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "invalid request body"})
		}

		token, err := useCase.SignUp(c, requestBody.ToEntity())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": err.Error()})
		}

		c.JSON(http.StatusOK, token)

		//if err := json.NewDecoder(r.Body()).Decode(&requestBody); err != nil {
		//	jsonBody := serializer.ErrorResponse(int(status.BadRequest), err, nil)
		//	return indigo.ErrorResponse(r, status.BadRequest, jsonBody)
		//}
		//
		//validator := validator.New()
		//if err := validator.Struct(requestBody); err != nil {
		//	jsonBody := serializer.ErrorResponse(int(status.BadRequest), err, nil)
		//	return indigo.ErrorResponse(r, status.BadRequest, jsonBody)
		//}
		//
		//token, err := useCase.SignUp(r.Ctx, requestBody.ToEntity())
		//if err != nil {
		//	jsonBody := serializer.ErrorResponse(int(status.InternalServerError), err, nil)
		//	return indigo.ErrorResponse(r, status.BadRequest, jsonBody)
		//}
		//
		//jsonBody := serializer.SuccessResponse(token)
		//return indigo.SuccessResponse(r, jsonBody)
	}
}
