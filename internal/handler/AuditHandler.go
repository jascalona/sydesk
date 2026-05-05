package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/audit"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuditHandler struct {
	Service audit.AuditService
}

func NewAuditHandler(s audit.AuditService) *AuditHandler {
	return &AuditHandler{Service: s}
}

func (h *AuditHandler) GetAudit(c *gin.Context) {
	parentStr := c.Query("parent_id")
	if parentStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parent_id es obligatorio"})
		return
	}
	parentID, err := uuid.Parse(parentStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parent_id debe ser un UUID válido"})
		return
	}

	audit, err := h.Service.GetAuditCustomer(c.Request.Context(), parentID)
	if err != nil {
		log.Printf("Error en el servicio %v:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusOK, audit)
}
