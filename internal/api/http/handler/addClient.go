package handler

import (
	"net/http"

	"github.com/drizzleent/vortex/internal/api"
	"github.com/gin-gonic/gin"
)

func (h *handler) AddClient(c *gin.Context) {
	h.s.AddClient()
	api.NewErrorResponse(c, http.StatusOK, "add")
}
