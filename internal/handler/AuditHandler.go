package handler

import (
	"log"
	"net/http"
	"strconv"
	"sydesk/internal/service/audit"

	"github.com/gin-gonic/gin"
)

type AuditHandler struct {
	Service audit.AuditService
}

func NewAuditHandler(s audit.AuditService) *AuditHandler {
	return &AuditHandler{Service: s}
}

func (h *AuditHandler) GetAudit(c *gin.Context) {
	customerID, err := strconv.Atoi(c.DefaultQuery("customer_id", "0"))
	if err != nil {
		log.Printf("Error al convertir customer_id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer_id invalido"})
		return
	}

	audit, err := h.Service.GetAuditCustomer(c.Request.Context(), customerID)
	if err != nil {
		log.Printf("Error en el servicio %v:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
	}
	c.JSON(http.StatusOK, audit)
}
