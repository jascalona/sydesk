package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/business"
	domain "sydesk/pkg/domain/business"
	"sydesk/pkg/utils"

	"github.com/gin-gonic/gin"
)

type TaskRequestHandler struct {
	Service business.TaskRequestServ
}

func NewTaskRequestHandler(s business.TaskRequestServ) *TaskRequestHandler {
	return &TaskRequestHandler{Service: s}
}

func (h *TaskRequestHandler) GetVWTaskRequest(c *gin.Context) {

	list_task, err := h.Service.GetVWTaskRequest(c.Request.Context())
	if err != nil {
		log.Println("Error al obtener los registros", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, list_task)

}

func (h *TaskRequestHandler) GetVWTaskRequestByTicket(c *gin.Context) {

	// Extraccion del ID
	ticket_id := c.Query("id")
	if ticket_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID de la solicitud es requerido"})
		return
	}
	taskById, err := h.Service.GetVWTaskRequestByTicket(c.Request.Context(), ticket_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: ": "No se pudo obtener los registros asociados"})
		return
	}
	c.JSON(http.StatusOK, taskById)
}

func (h *TaskRequestHandler) CreatedRequestTask(c *gin.Context) {
	var task domain.ValidateRequestTask

	if err := c.ShouldBindJSON(&task); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("Error en la validacion  del mensaje: ", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}

		log.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error:": "json mal formado"})
		return
	}

	// verificacion de campos
	validation_task := domain.TaskRequest{
		TICKET_ID:     task.TICKET_ID,
		PRIORITY:      task.PRIORITY,
		STATUS_ID:     task.STATUS_ID,
		CREATED_BY:    task.CREATED_BY,
		ASSIGNED_DPT:  task.ASSIGNED_DPT,
		ASSIGNED_USER: task.ASSIGNED_USER,
		TOPIC:         task.TOPIC,
		DESCRIPTION:   task.DESCRIPTION,
		EXPIRED_IN:    task.EXPIRED_IN,
		CLOSED_IN:     task.CLOSED_IN,
	}

	err := h.Service.CreatedRequestTask(c.Request.Context(), &validation_task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Solicitud procesada")
}
