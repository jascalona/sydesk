package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/organization"
	"sydesk/pkg/domain"

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
	var reqRole domain.Roles
	if err := c.ShouldBindJSON(&reqRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error al deserealizar el mensaje": err.Error()})
		return
	}

	err := h.Servcie.Created(c.Request.Context(), &reqRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errir interno": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Solicitud procesada")
}
