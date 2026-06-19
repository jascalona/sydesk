package router

import (
	"sydesk/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterAudit struct {
	audit_h         *handler.AuditHandler
	audit_service_h *handler.AsHandler
	channel_h       *handler.ChannelHandler
	item_h          *handler.ItemHandler
	sub_product_h   *handler.SupHandler
}

func NewRouterAudit(
	audit *handler.AuditHandler,
	audit_service *handler.AsHandler,
	channel *handler.ChannelHandler,
	item *handler.ItemHandler,
	sub_product *handler.SupHandler,
) *RouterAudit {
	return &RouterAudit{
		audit_h:         audit,
		audit_service_h: audit_service,
		channel_h:       channel,
		item_h:          item,
		sub_product_h:   sub_product,
	}
}

func (r *RouterAudit) RegisterAudit(rg *gin.RouterGroup) {

	audit := rg.Group("customer_audit")
	{
		audit.GET("", r.audit_h.GetAudit)
	}

	audit_service := rg.Group("tradetraceability")
	{
		audit_service.GET("", r.audit_service_h.AuditStatus)
		audit_service.POST("", r.audit_service_h.CreatedAS)
		audit_service.PATCH("", r.audit_service_h.Update)
	}

	channel := rg.Group("channel")
	{
		channel.GET("", r.channel_h.GetChannel)
	}

	item := rg.Group("items")
	{
		item.GET("", r.item_h.GetItem)
		item.POST("", r.item_h.CreatedItem)
	}

	sub_product := rg.Group("subproduct")
	{
		sub_product.GET("", r.sub_product_h.GetSup)
	}
}
