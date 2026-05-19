package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/audit"
	validation "sydesk/pkg/domain/audit"
	"sydesk/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	Service audit.ItemServ
}

func NewItemHandler(s audit.ItemServ) *ItemHandler {
	return &ItemHandler{Service: s}
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	items, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Printf("Error en el servicio %v", err.Error())
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) CreatedItem(c *gin.Context) {
	var itemsValidations validation.ValidationItem

	// validacion de entrada
	if err := c.ShouldBindJSON(&itemsValidations); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("error en la validacion del mensaje ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error de formato ": errors})
			return
		}
	}

	validate := validation.ItemActivities{
		PRODCUT_ID: *itemsValidations.PRODUCT_ID,
		NAME:       *itemsValidations.NAME,
	}

	// llamado del servicio
	err := h.Service.Created(c.Request.Context(), &validate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, "Solicitud procesada")

}
