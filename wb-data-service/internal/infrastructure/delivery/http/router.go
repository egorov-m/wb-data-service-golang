package http

import (
	"github.com/gin-gonic/gin"
	healthHandler "wb-data-service-golang/wb-data-service/internal/infrastructure/delivery/http/internal/handler/health"
	auth "wb-data-service-golang/wb-data-service/internal/module/authorization/core"
	authHandler "wb-data-service-golang/wb-data-service/internal/module/authorization/delivery/http/handler"
	"wb-data-service-golang/wb-data-service/internal/module/authorization/middleware"
	priceHistory "wb-data-service-golang/wb-data-service/internal/module/price-history/core"
	priceHistoryHandler "wb-data-service-golang/wb-data-service/internal/module/price-history/delivery/http/handler"
	product "wb-data-service-golang/wb-data-service/internal/module/product/core"
	productHandler "wb-data-service-golang/wb-data-service/internal/module/product/delivery/http/handler"
)

func InitRoutes(
	authModule auth.AuthorizationUseCase,
	productUseCase product.ProductUseCase,
	priceHistoryUseCase priceHistory.PriceHistoryUseCase,
) *gin.Engine {
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

		productGroup := apiV1Group.Group("/product", middleware.NewAuthMiddleware(authModule))
		{
			productGroup.POST("/load", productHandler.NewLoad(productUseCase))
			productGroup.GET("", productHandler.NewGetByNmId(productUseCase))
			productGroup.GET("/all", productHandler.NewGetAll(productUseCase))
			productGroup.GET("/count", productHandler.NewGetCount(productUseCase))
			productGroup.GET("/quantity", productHandler.NewGetQuantity(productUseCase))
			productGroup.DELETE("", productHandler.NewDeleteByNmId(productUseCase))
		}

		priceHistoryGroup := apiV1Group.Group("/price-history", middleware.NewAuthMiddleware(authModule))
		{
			priceHistoryGroup.POST("/load", priceHistoryHandler.NewLoad(priceHistoryUseCase))
			priceHistoryGroup.GET("", priceHistoryHandler.NewGetByProductNmId(priceHistoryUseCase))
		}
	}

	return router
}
