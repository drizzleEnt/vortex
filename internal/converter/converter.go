package converter

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/drizzleent/vortex/internal/model"
	"github.com/gin-gonic/gin"
)

const (
	id = "id"
)

func FromRequestToModel(c *gin.Context) (*model.Client, int, error) {
	var client model.Client

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = json.Unmarshal(body, &client)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &client, http.StatusOK, nil
}

func FromRequestToID(c *gin.Context) (int, int, error) {
	key := c.Param(id)
	if len(key) == 0 {
		return 0, http.StatusBadRequest, errors.New("client id is requared")
	}

	id, err := strconv.Atoi(key)
	if err != nil {
		return 0, http.StatusBadRequest, errors.New("client id is requared")
	}
	return id, http.StatusOK, nil
}

func FromRequestToAlgorithms(c *gin.Context) (*model.Algorithms, int, error) {
	var algos model.Algorithms

	err := c.Bind(&algos)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &algos, http.StatusOK, nil
}
