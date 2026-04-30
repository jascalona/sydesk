package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/organization"
	"sydesk/pkg/domain"
	"sydesk/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RolesHandler struct {
	Servcie organization.RolesServices
}

func NewRolesHandler(s organization.RolesServices) *RolesHandler {
	return &RolesHandler{Servcie: s}
}

func (h *RolesHandler) GetRoles(c *gin.Context) {

	roles, err := h.Servcie.GetAll(c.Request.Context())
	if err != nil {
		log.Printf("erro en el servicio", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func (h *RolesHandler) CreateRole(c *gin.Context) {
	// struct con el tags binding definido en la construccion del msj
	var reqRole domain.ValidateRoles

	// Al falla el bindeo por label o formato incorrecto
	if err := c.ShouldBindJSON(&reqRole); err != nil {
		// Retorno de msj de errores globales
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Printf("error en la validacion del mensaje", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON MAL FORMADO"})
		return
	}

	// caso de aprobacion
	role := domain.Roles{NAME: reqRole.NAME}

	if err := h.Servcie.Created(c.Request.Context(), &role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada")
}
