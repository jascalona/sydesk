package router

import (
	"sydesk/internal/handler"

	"github.com/gin-gonic/gin"
)

func ConfigUserRouter(r gin.IRouter, h *handler.UserHandler) {
	userRouter := r.Group("/users")
	{
		// Corregidos los nombres de los métodos para que coincidan con handler.go
		userRouter.POST("", h.CreateUser)
		userRouter.GET("", h.GetAllUsers)
	}
}
