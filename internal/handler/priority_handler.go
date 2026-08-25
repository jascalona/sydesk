package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/components"

	"github.com/gin-gonic/gin"
)

type PriorityHandler struct {
	Service components.PriorityServ
}

func NewPriorityHandler(s components.PriorityServ) *PriorityHandler {
	return &PriorityHandler{Service: s}
}

func (h *PriorityHandler) GetAll(c *gin.Context) {
	list_priority, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Println("Error al obtener los registros: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, list_priority)
}
