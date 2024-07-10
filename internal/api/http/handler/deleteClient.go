package handler

import (
	"net/http"

	"github.com/drizzleent/vortex/internal/api"
	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteClient(c *gin.Context) {
	h.s.DeleteClient()
	api.NewErrorResponse(c, http.StatusOK, "del")
}
