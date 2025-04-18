package handlers

import (
	"AvitoShop/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) BuyItem(c echo.Context) error {
	buyerUsername := c.QueryParam("buyer")
	item := c.Param("item")

	if buyerUsername == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "buyer username is required"})
	}

	if item == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "item is required"})
	}

	if err := services.BuyItem(h.DB, buyerUsername, item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}
