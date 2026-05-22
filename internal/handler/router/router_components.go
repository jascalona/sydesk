package router

import (
	"sydesk/internal/handler"

	"github.com/gin-gonic/gin"
)

// Agrupacion de hanlders
type RouterComponent struct {
	component_h             *handler.ComponentsHandler
	customer_role_h         *handler.CustomRoleHandler
	customer_product_role_h *handler.CustomerProductRoleHandler
	enviroment_h            *handler.EnviromentHandler
	product_h               *handler.ProductHandler
	status_h                *handler.StatusHandler
	subcomponent_h          *handler.SubcomponentHandler

	//priority_h *handler.P
}

// constructor de rutas para components
func NewRouterComponent(
	component *handler.ComponentsHandler,
	customer_role *handler.CustomRoleHandler,
	customer_product_role *handler.CustomerProductRoleHandler,
	enviroment *handler.EnviromentHandler,
	product *handler.ProductHandler,
	status *handler.StatusHandler,
	subcomponent *handler.SubcomponentHandler,
) *RouterComponent {
	return &RouterComponent{
		component_h:             component,
		customer_role_h:         customer_role,
		customer_product_role_h: customer_product_role,
		enviroment_h:            enviroment,
		product_h:               product,
		status_h:                status,
		subcomponent_h:          subcomponent,
	}
}

// construccion de los sub_grupos y endpoints en el enrutador principal
func (r *RouterComponent) RegisterComponents(rg *gin.RouterGroup) {

	// grupo de rutas por segmento

	// COMPONENTS
	components := rg.Group("product_components")
	{
		components.GET("", r.component_h.GetAllComponents)
		components.POST("", r.component_h.CreatedComponents)
	}

	customer_role := rg.Group("customers_roles")
	{
		customer_role.GET("", r.customer_role_h.GetAll)
	}

	customer_product_role := rg.Group("customers_product_roles")
	{
		customer_product_role.GET("", r.customer_product_role_h.GetAllCustomPR)
		customer_product_role.POST("", r.customer_product_role_h.CreatedCPR)
	}

	enviroments := rg.Group("enviroments")
	{
		enviroments.GET("", r.enviroment_h.GetEnviroments)
	}

	products := rg.Group("/products")
	{
		products.GET("", r.product_h.GetAllProduct)
		products.POST("", r.product_h.CreatedProduct)
	}

	status := rg.Group("status")
	{
		status.GET("", r.status_h.GetStatus)
	}

	subcomponent := rg.Group("subcomponents")
	{
		subcomponent.GET("", r.subcomponent_h.GetAll)
	}

}
