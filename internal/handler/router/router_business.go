package router

import (
	"sydesk/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterBusiness struct {
	customer_h *handler.CustomerHandler
	contact_h  *handler.ContactHandler

	ticket_h   *handler.TicketRequestHandler
	ticket_b_h *handler.TicketBreakHandler
}

func NewRouterBusiness(
	customer *handler.CustomerHandler,
	contact *handler.ContactHandler,
	ticket *handler.TicketRequestHandler,
	ticketBreak *handler.TicketBreakHandler,
) *RouterBusiness {
	return &RouterBusiness{
		customer_h: customer,
		contact_h:  contact,
		ticket_h:   ticket,
		ticket_b_h: ticketBreak,
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

	ticket := rg.Group("ticketrequest")
	{
		ticket.GET("", r.ticket_h.GetTicket)
		ticket.POST("", r.ticket_h.CreatedTicket)
	}

	ticket_b_h := rg.Group("ticketbreak")
	{
		ticket_b_h.GET("", r.ticket_b_h.GetTicketBreak)
		ticket_b_h.POST("", r.ticket_b_h.CreatedTBreak)

	}

}
