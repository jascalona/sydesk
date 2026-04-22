package router

import (
	"sydesk/internal/handler"
	"sydesk/pkg/domain"

	"github.com/gin-gonic/gin"
)

// SetupRouter es el administrador central de todos los endpoints de la API
func SetupRouter(r *gin.Engine, authService domain.AuthServices,
	userH *handler.UserHandler,
	productH *handler.ProductHandler,
	componentH *handler.ComponentsHandler,
) {

	// GESTION DE ACCESO PUBLICO (LOGIN Y "REGISTER SI APLICA")
	authH := &handler.AuthHandler{Service: authService}

	r.POST("/login", authH.Login)
	r.GET("/health", func(c *gin.Context) { c.Status(200) })

	// GESTION DE SERVICIOS PROTEGIDOS (API V1)
	// Grupo principal y aplicacion del bloqueo global
	api := r.Group("/api/v1")
	api.Use(handler.Auth(authService))

	{
		// --- GRUPO: USUARIOS ---
		users := api.Group("/users")
		{
			users.GET("", userH.GetAllUsers)
			users.POST("", userH.CreateUser)
		}

		// --- GRUPO: PRODUCTOS
		products := api.Group("/products")
		{
			products.GET("", productH.GetAllProduct)
			products.POST("", productH.CreatedProduct)
		}

		// GRUPO: COMPONENTS
		components := api.Group("/components")
		{
			components.GET("", componentH.GetAllComponents)
			components.POST("", componentH.CreatedComponents)
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
