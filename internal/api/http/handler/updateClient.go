package handler

import (
	"net/http"

	"github.com/drizzleent/vortex/internal/api"
	"github.com/drizzleent/vortex/internal/converter"
	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateClient(c *gin.Context) {
	ID, status, err := converter.FromRequestToID(c)
	if err != nil {
		api.NewErrorResponse(c, status, err.Error())
	}

	req, status, err := converter.FromRequestToModel(c)
	if err != nil {
		api.NewErrorResponse(c, status, err.Error())
	}

	err = h.s.UpdateClient(c, int64(ID), req)
	if err != nil {
		api.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(status, "clients client updated")
}
