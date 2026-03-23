package router

import (
	"sydesk/internal/handler"

	"github.com/gin-gonic/gin"
)

func ConfigUserRouter(r gin.IRouter, h *handler.UserHandler) {
	userRouter := r.Group("/users")
	{
		userRouter.POST("", h.CreateUser)
		userRouter.GET("", h.GetAllUsers)
	}
}
