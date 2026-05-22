package router

import (
	"sydesk/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterBusiness struct {
	customer_h *handler.CustomerHandler
}

func NewRouterBusiness(
	customer *handler.CustomerHandler,
) *RouterBusiness {
	return &RouterBusiness{
		customer_h: customer,
	}
}

func (r *RouterBusiness) RegisterBusiness(rg *gin.RouterGroup) {

	customers := rg.Group("customers")
	{
		customers.GET("", r.customer_h.GetCustomer)
		customers.POST("", r.customer_h.CreateCustomer)
	}
}
