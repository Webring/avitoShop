package handlers

import (
	"AvitoShop/internal/services"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) SendMoney(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	senderUsername, ok := claims["username"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Username not found in token")
	}

	receiverUsername := c.FormValue("toUser")
	valueStr := c.FormValue("amount")

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
