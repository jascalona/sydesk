package handler

import (
	"net/http"
	"sydesk/internal/service/components"

	"github.com/gin-gonic/gin"
)

type CustomRoleHandler struct {
	Service components.CustomRoleService
}

func NewCustomRoleHandler(s components.CustomRoleService) *CustomRoleHandler {
	return &CustomRoleHandler{Service: s}
}

func (h *CustomRoleHandler) GetAll(c *gin.Context) {
	customRele, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, customRele)
}
