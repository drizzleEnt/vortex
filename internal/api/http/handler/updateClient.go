package handler

import (
	"net/http"

	"github.com/drizzleent/vortex/internal/api"
	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateClient(c *gin.Context) {
	h.s.UpdateClient(c)
	api.NewErrorResponse(c, http.StatusOK, "updateClient")
}
