package handler

import (
	"net/http"
	"sydesk/internal/service/components"

	domain "sydesk/pkg/domain/components"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service components.ProductService
}

func NewProductHandler(s components.ProductService) *ProductHandler {
	return &ProductHandler{Service: s}
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	product, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		logger.Errorf("Error en el servicio GET product %v", err.Error())
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreatedProduct(c *gin.Context) {
	var reqProduct domain.Product
	if err := c.ShouldBindJSON(&reqProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error al deserializar el mensaje": err.Error()})
		return
	}

	err := h.Service.Created(c.Request.Context(), &reqProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada")
}
