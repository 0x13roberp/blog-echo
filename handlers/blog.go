package handlers

import (
	"blog/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetPost(c echo.Context) error {
	// TODO: logica para obtener los post del blog
	var post []models.Post

	if result := h.DB.Find(&post); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": result.Error.Error(),
		})
	}
	return c.JSON(http.StatusOK, post)
}

func (h *Handler) GetPostHtml(c echo.Context) error {
	var post []models.Post

	if result := h.DB.Find(&post); result.Error != nil {
		return c.String(http.StatusInternalServerError, "Error fetching posts") // error encontrando el post
	}

	// mostrar el post
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Posts": post,
	})
}
