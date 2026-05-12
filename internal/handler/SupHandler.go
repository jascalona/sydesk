package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/audit"

	"github.com/gin-gonic/gin"
)

type SupHandler struct {
	Service audit.SupService
}

func NewSupHandler(s audit.SupService) *SupHandler {
	return &SupHandler{Service: s}
}

func (h *SupHandler) GetSup(c *gin.Context) {
	subproduct, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Printf("Error en el servicio %v", err.Error())
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, subproduct)
}
