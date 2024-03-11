package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/core"
	"wb-data-service-golang/wb-data-service/internal/module/price-history/delivery/http/handler/request"
)

type PriceHistory interface {
	Load(context.Context, int) (core.PriceHistoryTask, error)
	GetByProductNmId(context.Context, int) ([]core.PriceHistory, error)
}

// @Summary  Load price history from wb
// @Tags     Price history
// @Accept   json
// @Param    requestBody         body request.LoadPriceHistoryInBody true "Nm id from load"
// @Success  202                 {object} core.PriceHistoryTask
// @Router   /price-history/load [post]
// @Security Bearer
func NewLoad(useCase core.PriceHistoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody request.LoadPriceHistoryInBody

		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				gin.H{"message": "invalid request body"})

			return
		}

		res, err := useCase.Load(c, requestBody.ToEntity())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"message": err.Error()})

			return
		}

		c.JSON(http.StatusAccepted, res)
	}
}

// @Summary  Get price history by product nm id
// @Tags     Price history
// @Produce  json
// @Param    nm_id          query    int true "Wb product nm id"
// @Success  200            {object} []core.PriceHistory
// @Router   /price-history [get]
// @Security Bearer
func NewGetByProductNmId(useCase core.PriceHistoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		nmIdStr := c.Query("nm_id")
		nmId, err := strconv.Atoi(nmIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nm_id parameter"})
			return
		}

		res, err := useCase.GetByProductNmId(c, nmId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound,
				gin.H{"message": err.Error()})

			return
		}

		c.JSON(http.StatusOK, res)
	}
}
