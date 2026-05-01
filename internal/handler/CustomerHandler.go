package handler

import (
	"log"
	"net/http"
	"strconv"
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

// --- MODULO AUDITORIA ---//
func (h *CustomerHandler) GetAuditCustomer(c *gin.Context) {
	customerID, err := strconv.Atoi(c.DefaultQuery("customer_id", "0"))
	if err != nil {
		log.Printf("Error al convertir el customer_id: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer_id invalido"})
		return
	}

	customers, err := h.Service.GetAuditCustomer(c.Request.Context(), customerID)
	if err != nil {
		log.Printf("Error en el servicio %v:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
	}
	c.JSON(http.StatusOK, customers)
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
