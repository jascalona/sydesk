package router

import (
	"sydesk/internal/handler"
	domain "sydesk/pkg/domain/organization"

	"github.com/gin-gonic/gin"
)

// routers structs
type MainRouters struct {
	ComponetRouter     *RouterComponent
	BusinessRouter     *RouterBusiness
	OrganizationRouter *RouterOrganization
	AuditRouter        *RouterAudit
}

// Funcion de carga manual para depuracion
func RegisterUser(r *gin.Engine, ru *handler.UserHandler) {
	users := r.Group("register")
	{
		users.POST("", ru.CreateUser)
	}
}

// SetupRouter es el administrador central de todos los endpoints de la API
func SetupRouter(r *gin.Engine, authService domain.AuthServices, routers MainRouters) {

	// GESTION DE ACCESO PUBLICO (LOGIN)
	// hay que agregar logica para el registro y recuperacion de clave
	authH := &handler.AuthHandler{Service: authService}

	r.POST("/login", authH.Login)
	r.GET("/health", func(c *gin.Context) { c.Status(200) })

	// GESTION DE SERVICIOS PROTEGIDOS (API V1)
	// Grupo principal y aplicacion del bloqueo global

	api_v1 := r.Group("/api/v1")
	api_v1.Use(handler.Auth(authService))
	{
		routers.ComponetRouter.RegisterComponents(api_v1.Group("/component"))

		routers.OrganizationRouter.RegisterOrganization(api_v1.Group("/organization"))

		routers.BusinessRouter.RegisterBusiness(api_v1.Group("/business"))

		routers.AuditRouter.RegisterAudit(api_v1.Group("/audit"))

	}
}
