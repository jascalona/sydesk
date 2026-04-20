package router

import (
	"sydesk/internal/handler"
	"sydesk/pkg/domain"

	"github.com/gin-gonic/gin"
)

// SetupRouter es el administrador central de todos los endpoints de la API
func SetupRouter(r *gin.Engine, authService domain.AuthService, userH *handler.UserHandler) {

	// 1. GESTIÓN DE ACCESO PÚBLICO
	// Estos endpoints no requieren validación de identidad
	authH := &handler.AuthHandler{Service: authService}

	r.POST("/login", authH.Login)
	r.GET("/health", func(c *gin.Context) { c.Status(200) })

	// 2. GESTIÓN DE SERVICIOS PROTEGIDOS (API V1)
	// Creamos el grupo principal y aplicamos el bloqueo global
	api := r.Group("/api/v1")
	api.Use(handler.Auth(authService))

	{
		// --- GRUPO: USUARIOS ---
		users := api.Group("/users")
		{
			users.POST("", userH.CreateUser)
			users.GET("", userH.GetAllUsers)
			// users.GET("/:id", userH.GetUserByID)
		}

		// --- GRUPO: TICKETS (Ejemplo de cómo expandir) ---
		// tickets := api.Group("/tickets")
		// {
		//     tickets.GET("", ticketH.GetAll)
		//     tickets.POST("", ticketH.Create)
		// }

		// --- GRUPO: DEPARTAMENTOS ---
		// depts := api.Group("/departments")
		// {
		//     depts.GET("", deptH.ListAll)
		// }
	}
}
