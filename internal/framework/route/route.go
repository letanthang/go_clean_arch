package route

import (
	"app/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(productHandler *handler.ProductHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/products/:id", productHandler.GetProductByID)
	r.GET("/products", productHandler.ListProducts)
	return r
}
