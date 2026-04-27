package handler

import (
	"net/http"
	services "sydesk/internal/service/components"

	"github.com/gin-gonic/gin"
)

type CustomerProductRoleHandler struct {
	Services services.CustomProductRoleService
}

func NewCustomerProductRoleHandler(s services.CustomProductRoleService) *CustomerProductRoleHandler {
	return &CustomerProductRoleHandler{Services: s}
}

func (h *CustomerProductRoleHandler) GetAllCustomPR(c *gin.Context) {
	customPR, err := h.Services.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customPR)
}
