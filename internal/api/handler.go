package api

import "github.com/gin-gonic/gin"

type Handler interface {
	AddClient(*gin.Context)
	UpdateClient(*gin.Context)
	DeleteClient(*gin.Context)
	UpdateAlgorithmStatus(*gin.Context)
	InitRoutes() *gin.Engine
}
