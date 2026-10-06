package handler

import (
	"log"
	"net/http"
	"regexp"
	"strings"
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

	toNullable := func(v int) *int {
		if v == 0 {
			return nil
		}
		return &v
	}

	// verificacion de campos
	validation_ticket := domain.TicketRequest{
		//	ID:              ticket.ID,
		AFFECTED_C:      ticket.AFFECTED_C,
		TICKET_BCV:      ticket.TICKET_BCV,
		ENVIROMENT:      ticket.ENVIROMENT,
		PRODUCT:         ticket.PRODUCT,
		PRIORITY:        ticket.PRIORITY,
		TYPE_REQUEST:    ticket.TYPE_REQUEST,
		SLA_ID:          ticket.SLA_ID,
		COMPONENT_ID:    ticket.COMPONENT_ID,
		SUBCOMPONENT_ID: ticket.SUBCOMPONENT_ID,
		CONTACT_ID:      toNullable(*ticket.CONTACT_ID),
		TOPIC:           ticket.TOPIC,
		DESCRIPTION:     ticket.DESCRIPTION,
		CREATED_BY:      ticket.CREATED_BY,
		EXPIRED_IN:      ticket.EXPIRED_IN,
	}

	err := h.Service.Created(c.Request.Context(), &validation_ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada")

}

func (h *TicketRequestHandler) UpdateTicket(c *gin.Context) {

	var validIDPattern = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	idParam := strings.TrimSpace(c.Param("id"))

	if idParam == "" {
		log.Printf("El ID del ticket es requerido: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID del ticket es requerido"})
		return
	}

	if !validIDPattern.MatchString(idParam) {
		log.Printf("El ID del ticket no es válido: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID del ticket no es válido"})
		return
	}

	var ticket domain.ValidateTicketRequest
	if err := c.ShouldBindJSON(&ticket); err != nil {
		errors := utils.GetValidationError(err)
		if errors != nil {
			log.Println("Error en la validación del mensaje: ", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}
		log.Println("Error en JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "json mal formado"})
		return
	}

	var contactID *int
	if ticket.CONTACT_ID != nil && *ticket.CONTACT_ID != 0 {
		contactID = ticket.CONTACT_ID
	}

	// Mapeamos asignando el ID al struct
	updateTicket := domain.TicketRequest{
		ID:              idParam, // <-- Aquí va el ID para que no quede la variable huérfana
		AFFECTED_C:      ticket.AFFECTED_C,
		TICKET_BCV:      ticket.TICKET_BCV,
		ENVIROMENT:      ticket.ENVIROMENT,
		PRODUCT:         ticket.PRODUCT,
		PRIORITY:        ticket.PRIORITY,
		TYPE_REQUEST:    ticket.TYPE_REQUEST,
		SLA_ID:          ticket.SLA_ID,
		COMPONENT_ID:    ticket.COMPONENT_ID,
		SUBCOMPONENT_ID: ticket.SUBCOMPONENT_ID,
		CONTACT_ID:      contactID,
		TOPIC:           ticket.TOPIC,
		DESCRIPTION:     ticket.DESCRIPTION,
		CREATED_BY:      ticket.CREATED_BY,
		EXPIRED_IN:      ticket.EXPIRED_IN,
	}

	// Llamamos a la función pasando SOLO los 2 argumentos que espera la interfaz
	if err := h.Service.UpdateByTicketId(c.Request.Context(), &updateTicket); err != nil {
		log.Println("Error al actualizar el registro: ", err.Error())
		log.Println("MSJ COMPLETO: %s", &updateTicket)
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Registro actualizado correctamente"})
}
