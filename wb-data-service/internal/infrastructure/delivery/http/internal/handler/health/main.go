package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewGetHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	}
}
