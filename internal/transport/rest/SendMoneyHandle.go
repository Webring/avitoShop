package handlers

import (
	"AvitoShop/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) SendMoney(c echo.Context) error {
	senderIdStr := c.FormValue("fromUser")
	receiverIdStr := c.FormValue("toUser")
	valueStr := c.FormValue("amount")

	senderId64, err := strconv.ParseUint(senderIdStr, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid fromUser"})
	}
	receiverId64, err := strconv.ParseUint(receiverIdStr, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid toUser"})
	}
	value64, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid amount"})
	}

	senderId := uint(senderId64)
	receiverId := uint(receiverId64)
	value := uint(value64)

	if err := services.SendMoney(h.DB, senderId, receiverId, value); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}
