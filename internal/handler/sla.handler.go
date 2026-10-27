package handler

import (
	"net/http"
	"sydesk/internal/service/components"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
)

type SLAHandler struct {
	Service components.SLAService
}

func NewSLAHandler(s components.SLAService) *SLAHandler {
	return &SLAHandler{Service: s}
}

func (h *SLAHandler) GetAllSLA(c *gin.Context) {
	slas, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		logger.Errorf("Error en el servicio GET sla %v", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, slas)

}
