package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/components"

	"github.com/gin-gonic/gin"
)

type TypeRequestHandler struct {
	Service components.TypeRequestServ
}

func NewTypeRequestHandler(s components.TypeRequestServ) *TypeRequestHandler {
	return &TypeRequestHandler{Service: s}
}

func (h *TypeRequestHandler) GetTypeRequest(c *gin.Context) {
	type_request, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Println("Error al obtener los registros ", err.Error())
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, type_request)
}
