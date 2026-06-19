package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/business"
	validation "sydesk/pkg/domain/business"
	"sydesk/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	Service business.ContactService
}

func NewContactHandler(s business.ContactService) *ContactHandler {
	return &ContactHandler{Service: s}
}

func (h *ContactHandler) GetContact(c *gin.Context) {

	contact, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) CreatedContact(c *gin.Context) {

	var reqContact validation.ContactValidation

	if err := c.ShouldBindJSON(&reqContact); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("error en la validacion del mensaje", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error: json mal formado": err.Error()})
		return
	}

	// validacion de la estrcutura
	contact := validation.Contact{
		CUSTOMER_ID: reqContact.CUSTOMER_ID,
		NAME:        reqContact.NAME,
		SURNAME:     reqContact.SURNAME,
		EMAIL:       reqContact.EMAIL,
		PHONE:       reqContact.PHONE,
	}

	if err := h.Service.Created(c.Request.Context(), &contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Solicitud procesada")
}
