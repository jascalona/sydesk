package handler

import (
	"log"
	"net/http"
	"strconv"
	"sydesk/internal/service/audit"
	validation "sydesk/pkg/domain/audit"
	"sydesk/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AsHandler struct {
	Service audit.AsService
}

func NewAsHandler(s audit.AsService) *AsHandler {
	return &AsHandler{Service: s}
}

func (h *AsHandler) AuditStatus(c *gin.Context) {

	parentStr := c.Query("cpr_id")
	if parentStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cpr_id es obligatorio"})
		return
	}
	cprID, err := uuid.Parse(parentStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parent_id debe ser un UUID válido"})
		return
	}

	as, err := h.Service.GetAll(c.Request.Context(), cprID)
	if err != nil {
		log.Printf("error interno: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, as)
}

func (h *AsHandler) CreatedAS(c *gin.Context) {
	// validacion de la construccion del msj (struct)
	var reqAS validation.ValidateAS

	if err := c.ShouldBindJSON(&reqAS); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Printf("error en la validacion del mensaje: %v", err)
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON MAL FORMADO"})
		return
	}

	toNullable := func(v string) *string {
		if v == "" {
			return nil
		}
		return &v
	}

	// aplicado de validaciones
	services := reqAS.SERVICES
	as := validation.AuditServ{
		CPR_ID:         reqAS.CPR_ID,
		START_AT:       toNullable(reqAS.START_AT),
		ENVIRONMENT:    reqAS.ENVIRONMENT,
		SERVICES:       &services,
		ACTIVITIES:     reqAS.ACTIVITIES,
		DESCRIPTION:    reqAS.DESCRIPTION,
		END_AT:         toNullable(reqAS.END_AT),
		LAST_OPERATION: toNullable(reqAS.LAST_OPERATION),
	}

	if err := h.Service.Created(c.Request.Context(), &as); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud Procesada")
}

func (h *AsHandler) Update(c *gin.Context) {
	idQuery := c.Query("id")
	id, err := strconv.ParseInt(idQuery, 10, 64)

	if err != nil {
		log.Printf("error al deserealizar el mensaje", err)
		c.JSON(http.StatusBadRequest, gin.H{"json mal formado": err})
		return
	}

	var input validation.ValidateAS

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Printf("error en la validacion del mensaje: %v", err)
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}
	}

	// llamada al servicio
	if err := h.Service.UpdateAS(c.Request.Context(), id, input); err != nil {
		log.Printf("Error al actualizar: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el registro"})
		return
	}
}
