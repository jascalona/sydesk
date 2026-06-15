package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/business"
	domain "sydesk/pkg/domain/business"
	"sydesk/pkg/utils"

	"github.com/gin-gonic/gin"
)

type TicketBreakHandler struct {
	Service business.TicketBreakServ
}

func NewTicketBreakHandler(s business.TicketBreakServ) *TicketBreakHandler {
	return &TicketBreakHandler{Service: s}
}

func (h *TicketBreakHandler) GetTicketBreak(c *gin.Context) {
	tBreak, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Println("Error al obtener los registros")
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, tBreak)
}

func (h *TicketBreakHandler) CreatedTBreak(c *gin.Context) {
	var t_break domain.ValidateTicketBreak

	if err := c.ShouldBindJSON(&t_break); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("Error en la validacion del mensaje ", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}

		log.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "json mal formado"})
		return
	}

	toNull := func(v string) *string {
		if v == "" {
			return nil
		}
		return &v
	}

	// verificacion de campos
	validation_ticket_break := domain.TicketBreak{
		TICKET_ID: t_break.TICKET_ID,
		STS_P:     t_break.STS_P,
		START_AT:  toNull(t_break.START_AT),
		END_AT:    toNull(t_break.END_AT),
	}

	err := h.Service.Created(c.Request.Context(), &validation_ticket_break)
	if err != nil {
		log.Println(validation_ticket_break)
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}

	log.Println(validation_ticket_break)
	c.JSON(http.StatusCreated, "Solicitud procesada")

}
