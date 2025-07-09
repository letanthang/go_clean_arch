package handler

import (
	"app/internal/domain/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	getUsecase  usecase.ProductGetUsecase
	listUsecase usecase.ProductListUsecase
}

func NewProductHandler(getUsecase usecase.ProductGetUsecase, listUsecase usecase.ProductListUsecase) *ProductHandler {
	return &ProductHandler{
		getUsecase:  getUsecase,
		listUsecase: listUsecase,
	}
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	product, err := h.getUsecase.Do(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.listUsecase.Do(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
