package handler

import (
	"net/http"
	"sydesk/internal/service/components"

	"github.com/gin-gonic/gin"
)

type EnviromentHandler struct {
	Service components.EnviromentServ
}

func NewEnviromentHandler(s components.EnviromentServ) *EnviromentHandler {
	return &EnviromentHandler{Service: s}
}

func (h *EnviromentHandler) GetEnviroments(c *gin.Context) {
	enviroment, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, enviroment)
}
