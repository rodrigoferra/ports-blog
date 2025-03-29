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

func (h *PostHandler) Create(ctx echo.Context) error {
	var dto domain.PostCreateDTO
	if err := ctx.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate input
	if err := ctx.Validate(dto); err != nil {
		return err
	}
	post, err := h.service.Create(ctx.Request().Context(), dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetByID(ctx echo.Context) error {
	ID := ctx.Param("id")
	post, err := h.service.GetByID(ctx.Request().Context(), ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if post == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}
	return ctx.JSON(http.StatusOK, post)
}

func (h *PostHandler) List(ctx echo.Context) error {
	posts, err := h.service.List(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, posts)
}

func (h *PostHandler) Update(ctx echo.Context) error {
	ID := ctx.Param("id")
	var dto domain.PostUpdateDTO
	if err := ctx.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate input
	if err := ctx.Validate(dto); err != nil {
		return err
	}
	if err := h.service.Update(ctx.Request().Context(), ID, dto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (h *PostHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")
	if err := h.service.Delete(ctx.Request().Context(), ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
