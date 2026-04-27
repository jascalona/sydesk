package handler

import (
	"net/http"
	"sydesk/internal/service/components"

	"github.com/gin-gonic/gin"
)

type SubcomponentHandler struct {
	Service components.SubcomponentService
}

func NewSubcomponentHandler(s components.SubcomponentService) *SubcomponentHandler {
	return &SubcomponentHandler{Service: s}
}

func (h *SubcomponentHandler) GetAll(c *gin.Context) {
	subcomponents, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subcomponents)
}
