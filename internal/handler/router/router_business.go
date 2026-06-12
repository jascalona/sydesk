package router

import (
	"sydesk/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterBusiness struct {
	customer_h *handler.CustomerHandler
	contact_h  *handler.ContactHandler

	ticket_h *handler.TicketRequestHandler
}

func NewRouterBusiness(
	customer *handler.CustomerHandler,
	contact *handler.ContactHandler,
	ticket *handler.TicketRequestHandler,
) *RouterBusiness {
	return &RouterBusiness{
		customer_h: customer,
		contact_h:  contact,
		ticket_h:   ticket,
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

	ticket := rg.Group("createdticket")
	{
		ticket.GET("", r.ticket_h.GetTicket)
		ticket.POST("", r.ticket_h.CreatedTicket)
	}
}
