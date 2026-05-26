package router

import (
	"sydesk/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterBusiness struct {
	customer_h *handler.CustomerHandler
	contact_h  *handler.ContactHandler
}

func NewRouterBusiness(
	customer *handler.CustomerHandler,
	contact *handler.ContactHandler,
) *RouterBusiness {
	return &RouterBusiness{
		customer_h: customer,
		contact_h:  contact,
	}
}

func (r *RouterBusiness) RegisterBusiness(rg *gin.RouterGroup) {

	customers := rg.Group("customers")
	{
		customers.GET("", r.customer_h.GetCustomer)
		customers.POST("", r.customer_h.CreateCustomer)
	}

	contact := rg.Group("contact")
	{
		contact.GET("", r.contact_h.GetContact)
		contact.POST("", r.contact_h.CreatedContact)
	}

}
