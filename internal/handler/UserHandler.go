package handler

import (
	"net/http"
	"sydesk/internal/service/organization"
	"sydesk/pkg/domain"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service organization.UserService
}

func NewUserHandler(s organization.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var reqUser domain.User
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error al deserializar el mensaje": err.Error()})
		return
	}

	err := h.Service.Create(c.Request.Context(), &reqUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error interno": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada")

}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
