package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
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

func (h *NotesHandler) UpdateNotes(c *gin.Context) {
	idParam := strings.TrimSpace(c.Param("id"))

	// validacion y conversion
	noteId, err := strconv.Atoi(idParam)
	if err != nil || noteId <= 0 {
		log.Printf("El ID no es valido: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "El ID debe ser un numero valido"})
		return
	}

	var note domain.ValidateNotes
	if err := c.ShouldBindJSON(&note); err != nil {
		errors := utils.GetValidationError(err)
		if errors != nil {
			log.Println("Error en la validacion del mensaje: ", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error de formato: ": errors})
			return
		}
		log.Println("Error en el JSON", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "json mal formado"})
		return
	}

	validation_notes := domain.Notes{
		ID:          noteId,
		TASK_ID:     note.TASK_ID,
		TOPIC:       note.TOPIC,
		DESCRIPTION: note.DESCRIPTION,
		CREATED_BY:  note.CREATED_BY,
	}

	if err := h.Service.UpdateNotes(c.Request.Context(), &validation_notes); err != nil {
		log.Println("Error al actualizar el registro: ", err.Error())
		log.Printf("MSJ COMPLETO: %+v", validation_notes)
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Solicitud Procesada")

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
