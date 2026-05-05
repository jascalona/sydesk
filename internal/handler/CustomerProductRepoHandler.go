package handler

import (
	"log"
	"net/http"
	services "sydesk/internal/service/components"
	validation "sydesk/pkg/domain/components"
	"sydesk/pkg/utils"

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

func (h *CustomerProductRoleHandler) CreatedCPR(c *gin.Context) {
	var reqCPR validation.ValidationCustomerProductRole

	if err := c.ShouldBindJSON(&reqCPR); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Printf("error en la validacion del mensaje: %v", err)
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}
		log.Printf("error al deserealizar el mensaje: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "json mal formado"})
		return
	}

	// verificacion de campos
	custompr := validation.CustomerProductRole{
		CUSTOMER_ID: reqCPR.CUSTOMER_ID,
		PRODUCT_ID:  reqCPR.PRODUCT_ID,
		ROLE_ID:     reqCPR.ROLE_ID,
		PARENT_ID:   reqCPR.PARENT_ID.UUID,
	}

	err := h.Services.Created(c.Request.Context(), &custompr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada")
}
