package handlers

import (
	"AvitoShop/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) BuyItem(c echo.Context) error {
	buyerIdStr := c.QueryParam("buyer")
	item := c.Param("item")

	buyerId64, err := strconv.ParseUint(buyerIdStr, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid buyer id"})
	}

	buyerId := uint(buyerId64)

	if err := services.BuyItem(h.DB, buyerId, item); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}
