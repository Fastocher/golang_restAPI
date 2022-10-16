package handler

import (
	"github.com/Fastocher/restapp/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.singIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		messages := api.Group("/messages")
		{
			messages.GET("/")
			messages.GET("/:id")
			messages.POST("/", h.CreateMessage)
			messages.PUT("/:id")
			messages.DELETE("/:id")
		}

	}

	return router
}
