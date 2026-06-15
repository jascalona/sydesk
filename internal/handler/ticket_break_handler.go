package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/business"

	"github.com/gin-gonic/gin"
)

type TicketBreakHandler struct {
	Service business.TicketBreakServ
}

func NewTicketBreakHandler(s business.TicketBreakServ) *TicketBreakHandler {
	return &TicketBreakHandler{Service: s}
}

func (h *TicketBreakHandler) GetTicketBreak(c *gin.Context) {
	tBreak, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Println("Error al obtener los registros")
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, tBreak)
}
