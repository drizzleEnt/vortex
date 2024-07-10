package handler

import (
	"net/http"

	"github.com/drizzleent/vortex/internal/api"
	"github.com/drizzleent/vortex/internal/converter"
	"github.com/gin-gonic/gin"
)

func (h *handler) AddClient(c *gin.Context) {
	req, status, err := converter.FromRequestToModel(c)
	if err != nil {
		api.NewErrorResponse(c, status, err.Error())
	}
	resp, err := h.s.AddClient(c, req)
	if err != nil {
		api.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(status, resp)
}
