package handler

import (
	"net/http"

	"github.com/drizzleent/vortex/internal/api"
	"github.com/drizzleent/vortex/internal/converter"
	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteClient(c *gin.Context) {
	req, status, err := converter.FromRequestToID(c)
	if err != nil {
		api.NewErrorResponse(c, status, err.Error())
	}
	err = h.s.DeleteClient(c, int64(req))
	if err != nil {
		api.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(status, "client deleted")
}
