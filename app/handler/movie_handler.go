package handler

import (
	"github.com/labstack/echo"
	"github.com/sidie88/stockbit/app/service"
	"net/http"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

type MovieHandler struct {
	OutboundService *service.OmDbApiServiceImpl
}

func NewMovieHandler(s *service.OmDbApiServiceImpl) *MovieHandler{
	handler := &MovieHandler{
		OutboundService: s,
	}
	return handler
}

func (m *MovieHandler) SearchMovie(c echo.Context) error{
	sw := c.QueryParam("searchword")
	p := c.QueryParam("pagination")
	resp, err := m.OutboundService.SearchWithPagination(p, sw)
	if err != nil {
		return c.JSON(500, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (m *MovieHandler) GetMovieDetail(c echo.Context) error{
	imdbid := c.QueryParam("imdbid")
	resp, err := m.OutboundService.GetMovieDetail(imdbid)
	if err != nil {
		return c.JSON(500, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}