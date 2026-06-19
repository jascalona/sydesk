package router

import (
	"sydesk/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterOrganization struct {
	users_h *handler.UserHandler
	roles_h *handler.RolesHandler
}

func NewRouterOrganization(
	users *handler.UserHandler,
	roles *handler.RolesHandler,
) *RouterOrganization {
	return &RouterOrganization{
		users_h: users,
		roles_h: roles,
	}
}

func (r *RouterOrganization) RegisterOrganization(rg *gin.RouterGroup) {

	users := rg.Group("users")
	{
		users.GET("", r.users_h.GetAllUsers)
		users.POST("", r.users_h.CreateUser)
	}

	roles := rg.Group("roles")
	{
		roles.GET("", r.roles_h.GetRoles)
		roles.POST("", r.roles_h.CreateRole)
	}
}
