package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/business"
	domain "sydesk/pkg/domain/business"
	"sydesk/pkg/utils"

	"github.com/gin-gonic/gin"
)

type TicketRequestHandler struct {
	Service business.TicketRequestServ
}

func NewTicketRequestHandler(s business.TicketRequestServ) *TicketRequestHandler {
	return &TicketRequestHandler{Service: s}
}

func (h *TicketRequestHandler) GetVWTicketRequest(c *gin.Context) {
	vw_ticket, err := h.Service.GetVWTicketRequest(c.Request.Context())
	if err != nil {
		log.Println("Error al obtener los registros: ", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, vw_ticket)
}

func (h *TicketRequestHandler) GetTicket(c *gin.Context) {
	ticket, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		log.Println("Error al obtener los registros ", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (h *TicketRequestHandler) CreatedTicket(c *gin.Context) {
	var ticket domain.ValidateTicketRequest

	if err := c.ShouldBindJSON(&ticket); err != nil {
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

	// verificacion de campos
	validation_ticket := domain.TicketRequest{
		//	ID:              ticket.ID,
		AFFECTED_C:      ticket.AFFECTED_C,
		TICKET_BCV:      ticket.TICKET_BCV,
		ENVIROMENT:      ticket.ENVIROMENT,
		PRIORITY:        ticket.PRIORITY,
		TYPE_REQUEST:    ticket.TYPE_REQUEST,
		SLA_ID:          ticket.SLA_ID,
		COMPONENT_ID:    ticket.COMPONENT_ID,
		SUBCOMPONENT_ID: ticket.SUBCOMPONENT_ID,
		CONTACT_ID:      ticket.CONTACT_ID,
		TOPIC:           ticket.TOPIC,
		DESCRIPTION:     ticket.DESCRIPTION,
		CREATED_BY:      ticket.CREATED_BY,
		CREATED_AT:      ticket.CREATED_AT,
		EXPIRED_IN:      ticket.EXPIRED_IN,
	}

	err := h.Service.Created(c.Request.Context(), &validation_ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada")

}
