package handler

import (
	"log"
	"net/http"
	"sydesk/internal/service/organization"
	"sydesk/pkg/domain"
	"sydesk/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service organization.UserService
}

func NewUserHandler(s organization.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	// struct con el tags binding definido en la construccion del msj
	var reqUser domain.ValidationUser

	// al fallar el bindeo por label o formato incorrecto
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		// Retorno de msj de errores globales
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("error en la validacion del mensaje", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error de formato": errors})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "json mal formado"})
		return
	}

	// caso de aprobacion
	user := domain.User{
		ID:             reqUser.ID,
		ALIAS:          reqUser.ALIAS,
		NAME:           reqUser.NAME,
		SURNAME:        reqUser.SURNAME,
		EMAIL:          reqUser.EMAIL,
		PHONE:          reqUser.PHONE,
		DEPARTAMENT_ID: reqUser.DEPARTAMENT_ID,
		ROLE_ID:        reqUser.ROLE_ID,
		IS_ACTIVE:      reqUser.IS_ACTIVE,
	}

	if err := h.Service.Create(c.Request.Context(), &user); err != nil {
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
