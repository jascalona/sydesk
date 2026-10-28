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

	task_h  *handler.TaskRequestHandler
	notes_h *handler.NotesHandler
}

func NewRouterBusiness(
	customer *handler.CustomerHandler,
	contact *handler.ContactHandler,
	ticket *handler.TicketRequestHandler,
	ticketBreak *handler.TicketBreakHandler,
	taskRequest *handler.TaskRequestHandler,
	notes *handler.NotesHandler,
) *RouterBusiness {
	return &RouterBusiness{
		customer_h: customer,
		contact_h:  contact,
		ticket_h:   ticket,
		ticket_b_h: ticketBreak,
		task_h:     taskRequest,
		notes_h:    notes,
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
		ticket.PUT("/:id", r.ticket_h.UpdateTicket)
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
		task_request.PUT("/:id", r.task_h.UpdateRequestTask)
	}

	task_by_ticket := rg.Group("taskbyticket")
	{
		task_by_ticket.GET("", r.task_h.GetVWTaskRequestByTicket)
	}

	notes := rg.Group("notes")
	{
		notes.POST("", r.notes_h.CreatedNotes)
		notes.GET("", r.notes_h.GetByTaskId)
	}

}
