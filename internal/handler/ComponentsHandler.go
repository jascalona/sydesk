package handler

import (
	"net/http"
	services "sydesk/internal/service/components"
	components "sydesk/pkg/domain/components"

	"github.com/gin-gonic/gin"
)

type ComponentsHandler struct {
	Service services.ComponentService
}

func NewComponentsHandler(s services.ComponentService) *ComponentsHandler {
	return &ComponentsHandler{Service: s}
}

func (h *ComponentsHandler) GetAllComponents(c *gin.Context) {
	component, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, component)
}

func (h *ComponentsHandler) CreatedComponents(c *gin.Context) {
	var reqComp components.Components
	if err := c.ShouldBindJSON(&reqComp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error al deserealizar el mensaje": err.Error()})
		return
	}

	err := h.Service.Created(c.Request.Context(), &reqComp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada")
}
