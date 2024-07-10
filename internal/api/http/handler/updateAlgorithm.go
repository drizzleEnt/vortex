package handler

import (
	"net/http"

	"github.com/drizzleent/vortex/internal/api"
	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateAlgorithmStatus(c *gin.Context) {
	h.s.UpdateAlgorithmStatus(c)
	api.NewErrorResponse(c, http.StatusOK, "updateAlgo")
}
