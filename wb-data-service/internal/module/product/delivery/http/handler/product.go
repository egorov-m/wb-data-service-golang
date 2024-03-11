package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	product "wb-data-service-golang/wb-data-service/internal/module/product/core"
	"wb-data-service-golang/wb-data-service/internal/module/product/delivery/http/handler/request"
)

type Product interface {
	Load(ctx context.Context, product product.Product) (product.ProductTask, error)
	GetByNmId(ctx context.Context) (product.Product, error)
	GetAll(ctx context.Context) ([]product.Product, error)
	NewGetCount(ctx context.Context) (int, error)
	NewGetQuantity(ctx context.Context) (int, error)
	NewDeleteByNmId(ctx context.Context) (bool, error)
}

// @Summary  Load product from wb
// @Tags     Product
// @Accept   json
// @Param    requestBody   body     request.LoadProductInBody true "Nm id from load"
// @Success  202           {object} core.ProductTask
// @Router   /product/load [post]
// @Security Bearer
func NewLoad(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody request.LoadProductInBody

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

// @Summary  Get product by nm id
// @Tags     Product
// @Produce  json
// @Param    nm_id    query int true "Wb nm id"
// @Success  200      {object} core.Product
// @Router   /product [get]
// @Security Bearer
func NewGetByNmId(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		nmIdStr := c.Query("nm_id")
		nmId, err := strconv.Atoi(nmIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nm_id parameter"})
			return
		}

		res, err := useCase.GetByNmId(c, product.Product{NmId: nmId})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound,
				gin.H{"message": err.Error()})

			return
		}

		c.JSON(http.StatusOK, res)
	}
}

// @Summary  Get all products
// @Tags     Product
// @Success  200 {object} []core.Product
// @Router   /product/all [get]
// @Security Bearer
func NewGetAll(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := useCase.GetAll(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"message": err.Error()})

			return
		}

		c.JSON(http.StatusOK, res)
	}
}

// @Summary  Get count products
// @Tags     Product
// @Success  200            {object} int
// @Router   /product/count [get]
// @Security Bearer
func NewGetCount(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := useCase.GetCount(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"message": err.Error()})

			return
		}

		c.String(http.StatusOK, strconv.Itoa(res))
	}
}

// @Summary  Quantity of products available
// @Tags     Product
// @Success  200              {object} int
// @Router   /product/quantity [get]
// @Security Bearer
func NewGetQuantity(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := useCase.GetQuantity(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"message": err.Error()})

			return
		}

		c.String(http.StatusOK, strconv.Itoa(res))
	}
}

// @Summary  Delete product by nm id
// @Tags     Product
// @Param    nm_id    query    int true "Wb nm id"
// @Success  204      {object} bool
// @Router   /product [delete]
// @Security Bearer
func NewDeleteByNmId(useCase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		nmIdStr := c.Query("nm_id")
		nmId, err := strconv.Atoi(nmIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nm_id parameter"})
			return
		}

		err = useCase.DeleteByNmId(c, product.Product{NmId: nmId})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"message": err.Error()})

			return
		}

		c.String(http.StatusNoContent, "true")
	}
}
