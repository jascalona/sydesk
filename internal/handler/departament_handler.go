package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/components"

	"github.com/gin-gonic/gin"
)

type DepartamentHandler struct {
	Service components.DepartamentService
}

func NewDepartamentHandler(s components.DepartamentService) *DepartamentHandler {
	return &DepartamentHandler{Service: s}
}

func (h *DepartamentHandler) GetAll(c *gin.Context) {
	list_departaments, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Printf("Error al obtener los departamentos: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, list_departaments)
}
