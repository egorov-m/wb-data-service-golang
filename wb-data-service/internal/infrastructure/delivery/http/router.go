package http

import (
	"github.com/gin-gonic/gin"
	healthHandler "wb-data-service-golang/wb-data-service/internal/infrastructure/delivery/http/internal/handler/health"
	auth "wb-data-service-golang/wb-data-service/internal/module/authorization/core"
	authHandler "wb-data-service-golang/wb-data-service/internal/module/authorization/delivery/http/handler"
	"wb-data-service-golang/wb-data-service/internal/module/authorization/middleware"
)

func InitRoutes(authModule auth.AuthorizationUseCase) *gin.Engine {
	router := gin.New()

	healthGroup := router.Group("/health")
	{
		healthGroup.GET("", healthHandler.NewGetHealth())
	}

	apiV1Group := router.Group("api/v1/")
	{
		authGroup := apiV1Group.Group("/auth")
		{
			authGroup.POST("/sign-in", authHandler.NewSignIn(authModule))
			authGroup.POST("/sign-up", authHandler.NewSignUp(authModule))
		}
		mainGroup := apiV1Group.Group("", middleware.NewAuthMiddleware(authModule))
		{
			mainGroup.GET("/test")
		}
	}

	return router
}
