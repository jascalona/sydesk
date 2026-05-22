package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/audit"

	"github.com/gin-gonic/gin"
)

type ChannelHandler struct {
	Service audit.ChannelServ
}

func NewChannelHandler(s audit.ChannelServ) *ChannelHandler {
	return &ChannelHandler{Service: s}
}

func (h *ChannelHandler) GetChannel(c *gin.Context) {
	channel, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Printf("Error al obtener los registros %v", err.Error())
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, channel)
}
