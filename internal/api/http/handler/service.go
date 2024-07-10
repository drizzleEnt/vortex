package handler

import "github.com/gin-gonic/gin"

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/")
	{
		api.GET("/")
		api.POST("/create", h.AddClient)
		api.POST("/edit/:id", h.UpdateClient)
		api.POST("/delete/:id", h.DeleteClient)
		api.POST("/editAlgo/:id", h.UpdateAlgorithmStatus)

	}

	return router
}
