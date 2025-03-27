package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rodrigoferra/ports-blog/internal/core/domain"
	"github.com/rodrigoferra/ports-blog/internal/core/ports"
	"net/http"
)

type PostHandler struct {
	service ports.PostService
}

func NewPostHandler(service ports.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) Create(c echo.Context) error {
	var dto domain.PostCreateDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate input
	if err := c.Validate(dto); err != nil {
		return err
	}
	post, err := h.service.Create(c.Request().Context(), dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetByID(c echo.Context) error {
	ID := c.Param("id")
	post, err := h.service.GetByID(c.Request().Context(), ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if post == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}
	return c.JSON(http.StatusOK, post)
}

func (h *PostHandler) List(c echo.Context) error {
	posts, err := h.service.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) Update(c echo.Context) error {
	ID := c.Param("id")
	var dto domain.PostUpdateDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate input
	if err := c.Validate(dto); err != nil {
		return err
	}
	if err := h.service.Update(c.Request().Context(), ID, dto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *PostHandler) Delete(c echo.Context) error {
	ID := c.Param("id")
	if err := h.service.Delete(c.Request().Context(), ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
