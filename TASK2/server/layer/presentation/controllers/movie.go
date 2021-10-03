package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rampo0/bibit/layer/application/services"
	"github.com/rampo0/bibit/layer/domain"
	"github.com/rampo0/bibit/utils/errors"
)

type MovieHandler interface {
	Detail(c *gin.Context)
	Search(c *gin.Context)
}

type movieHandler struct {
	service services.MovieService
}

func NewMovieHandler(service services.MovieService) MovieHandler {
	return &movieHandler{service}
}

func (h *movieHandler) Detail(c *gin.Context) {
	var req domain.MovieDetailRequest
	err := c.ShouldBind(&req)
	if err != nil {
		err := errors.NewInternalServerError()
		c.JSON(err.Status, err)
		return
	}
	result, restErr := h.service.Detail(req.ID)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *movieHandler) Search(c *gin.Context) {
	var req domain.SearchMovieRequest
	err := c.ShouldBind(&req)
	if err != nil {
		err := errors.NewInternalServerError()
		c.JSON(err.Status, err)
		return
	}

	result, restErr := h.service.Search(req.Search, req.Page)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
