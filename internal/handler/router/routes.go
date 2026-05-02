package router

import (
	"sydesk/internal/handler"
	"sydesk/pkg/domain"

	"github.com/gin-gonic/gin"
)

// SetupRouter es el administrador central de todos los endpoints de la API
func SetupRouter(r *gin.Engine, authService domain.AuthServices,
	userH *handler.UserHandler,
	rolesH *handler.RolesHandler,

	productH *handler.ProductHandler,
	componentH *handler.ComponentsHandler,
	subcomponetH *handler.SubcomponentHandler,
	envirometH *handler.EnviromentHandler,
	statusH *handler.StatusHandler,
	customerH *handler.CustomerHandler,
	customReoleH *handler.CustomRoleHandler,
	customPRH *handler.CustomerProductRoleHandler,
	auditH *handler.AuditHandler,
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

		roles := api.Group("/rolesuser")
		{
			roles.GET("", rolesH.GetRoles)
			roles.POST("", rolesH.CreateRole)
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

		// GRUPO: Subcomponentes
		subcomponents := api.Group("/subcomponents")
		{
			subcomponents.GET("", subcomponetH.GetAll)
		}

		// GRUPO: Enviroment
		enviroment := api.Group("/enviroment")
		{
			enviroment.GET("", envirometH.GetEnviroments)
		}

		// GRUPO: STATUS
		status := api.Group("/status")
		{
			status.GET("", statusH.GetStatus)
		}

		// GRUPO: CUSTOMERS
		customer := api.Group("/customers")
		{
			customer.GET("", customerH.GetCustomer)
			customer.POST("", customerH.CreateCustomer)
		}

		// --- MODULO AUDITORIA ---//
		audit_service := api.Group("/auditcustom")
		{
			audit_service.GET("", auditH.GetAudit)
		}

		customRole := api.Group("/custom_roles")
		{
			customRole.GET("", customReoleH.GetAll)
		}

		customPR := api.Group("/customerpr")
		{
			customPR.GET("", customPRH.GetAllCustomPR)
		}

	}
}
