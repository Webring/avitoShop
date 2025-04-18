package handlers

import (
	"AvitoShop/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) SendMoney(c echo.Context) error {
	senderUsername := c.FormValue("fromUser")
	receiverUsername := c.FormValue("toUser")
	valueStr := c.FormValue("amount")

	if senderUsername == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "fromUser (username) is required"})
	}
	if receiverUsername == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "toUser (username) is required"})
	}
	if valueStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "amount is required"})
	}

	value64, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid amount"})
	}
	value := uint(value64)

	if err := services.SendMoney(h.DB, senderUsername, receiverUsername, value); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}
