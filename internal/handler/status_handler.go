package handler

import (
	"net/http"
	"sydesk/internal/service/components"

	"github.com/gin-gonic/gin"
)

type StatusHandler struct {
	Service components.StatusService
}

func NewStatusHandler(s components.StatusService) *StatusHandler {
	return &StatusHandler{Service: s}
}

func (h *StatusHandler) GetStatus(c *gin.Context) {
	status, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, status)
}
