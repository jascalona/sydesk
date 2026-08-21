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

	task_h *handler.TaskRequestHandler
}

func NewRouterBusiness(
	customer *handler.CustomerHandler,
	contact *handler.ContactHandler,
	ticket *handler.TicketRequestHandler,
	ticketBreak *handler.TicketBreakHandler,
	taskRequest *handler.TaskRequestHandler,
) *RouterBusiness {
	return &RouterBusiness{
		customer_h: customer,
		contact_h:  contact,
		ticket_h:   ticket,
		ticket_b_h: ticketBreak,
		task_h:     taskRequest,
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
		ticket.GET("", r.ticket_h.GetVWTicketRequest)
		ticket.POST("", r.ticket_h.CreatedTicket)
	}

	ticket_b_h := rg.Group("ticketbreak")
	{
		ticket_b_h.GET("", r.ticket_b_h.GetTicketBreak)
		ticket_b_h.POST("", r.ticket_b_h.CreatedTBreak)
		ticket_b_h.GET("/id/:id/", r.ticket_b_h.StatusTicketById)
	}

	task_request := rg.Group("taskrequest")
	{
		task_request.POST("", r.task_h.CreatedRequestTask)
		task_request.GET("", r.task_h.GetVWTaskRequest)
	}

	task_by_ticket := rg.Group("taskbyticket")
	{
		task_by_ticket.GET("", r.task_h.GetVWTaskRequestByTicket)
	}

}
