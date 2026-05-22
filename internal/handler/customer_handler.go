package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/business"
	domain "sydesk/pkg/domain/business"
	"sydesk/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	Service business.CustomerServ
}

func NewCustomerHandler(s business.CustomerServ) *CustomerHandler {
	return &CustomerHandler{Service: s}
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	customer, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Printf("Error en el servicio %v:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
	}
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var reqCustom domain.ValidateCustomer

	if err := c.ShouldBindJSON(&reqCustom); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("error en la validacion del mensaje", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "json mal formado"})
		return
	}

	toNullable := func(v string) *string {
		if v == "" {
			return nil
		}
		return &v
	}

	// verificacion de campos
	custom := domain.Customer{
		RIF:        reqCustom.RIF,
		NAME:       reqCustom.NAME,
		CHANNEL:    toNullable(reqCustom.CHANNEL),
		WS_GROUP:   toNullable(reqCustom.WS_GROUP),
		UID_SYPAGO: toNullable(reqCustom.UID_SYPAGO),
	}

	err := h.Service.Created(c.Request.Context(), &custom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada")

}
