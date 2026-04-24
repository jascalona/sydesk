package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/business"
	domain "sydesk/pkg/domain/business"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	Service business.CustomerServ
}

func NewCustomerHandler(s business.CustomerServ) *CustomerHandler {
	return &CustomerHandler{Service: s}
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	customer, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Printf("Error en el servicio %v:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
	}
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var reqCustom domain.Customer
	if err := c.ShouldBindJSON(&reqCustom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error al deserealizar el mensaje": err.Error()})
		return
	}

	err := h.Service.Created(c.Request.Context(), &reqCustom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada")

}
