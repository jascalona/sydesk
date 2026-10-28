package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/business"

	domain "sydesk/pkg/domain/business"
	"sydesk/pkg/utils"

	"github.com/gin-gonic/gin"
)

type NotesHandler struct {
	Service business.NotesService
}

func NewNotesHandler(s business.NotesService) *NotesHandler {
	return &NotesHandler{
		Service: s,
	}
}

func (h *NotesHandler) GetByTaskId(c *gin.Context) {

	// Extraccion del ID
	task_id := c.Query("task_id")
	if task_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID de la tarea es requerido"})
		return
	}

	taskById, err := h.Service.GetByTaskId(c.Request.Context(), task_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: ": "No se pudo obtener los registros asociados"})
		return
	}
	c.JSON(http.StatusOK, taskById)
}

func (h *NotesHandler) CreatedNotes(c *gin.Context) {
	var notes domain.ValidateNotes

	if err := c.ShouldBindJSON(&notes); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("Error en la validacion  del mensaje: ", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}

		log.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "json mal formado"})
		return
	}

	// validacion de campos
	validate_notes := domain.Notes{
		TASK_ID:     notes.TASK_ID,
		TOPIC:       notes.TOPIC,
		DESCRIPTION: notes.DESCRIPTION,
		CREATED_BY:  notes.CREATED_BY,
	}

	err := h.Service.CreatedNotes(c.Request.Context(), &validate_notes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno: ": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Solicitud procesada")

}
