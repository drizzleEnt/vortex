package handler

import (
	"github.com/drizzleent/vortex/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	s service.ApiService
}

func NewHandler(service service.ApiService) *handler {
	return &handler{
		s: service,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/")
	{
		api.POST("/create", h.AddClient)
		api.POST("/edit/:id", h.UpdateClient)
		api.POST("/delete/:id", h.DeleteClient)
		api.POST("/editAlgo/:id", h.UpdateAlgorithmStatus)

	}

	return router
}
