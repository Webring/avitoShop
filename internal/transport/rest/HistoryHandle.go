package handlers

import (
	"AvitoShop/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type MoneyTransactionDTO struct {
	Amount uint `json:"amount"`
	ID     uint `json:"id"`
}

type ReceivedMoneyTransactionDTO struct {
	MoneyTransactionDTO
	FromUser string `json:"fromUser"`
}

type SentMoneyTransactionDTO struct {
	MoneyTransactionDTO
	ToUser string `json:"toUser"`
}

func (h *Handler) MoneyHistory(c echo.Context) error {
	userIdStr := c.QueryParam("user")

	userId64, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid user id"})
	}

	userId := uint(userId64)

	sended, err := services.SendedMoney(h.DB, userId)

	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	received, err := services.RecievedMoney(h.DB, userId)

	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	sentDTOs := make([]SentMoneyTransactionDTO, 0, len(sended))
	for _, t := range sended {
		sentDTOs = append(sentDTOs, SentMoneyTransactionDTO{
			MoneyTransactionDTO: MoneyTransactionDTO{
				ID:     t.ID,
				Amount: t.Value,
			},
			ToUser: t.ToUser.Username,
		})
	}

	receivedDTOs := make([]ReceivedMoneyTransactionDTO, 0, len(received))
	for _, t := range received {
		receivedDTOs = append(receivedDTOs, ReceivedMoneyTransactionDTO{
			MoneyTransactionDTO: MoneyTransactionDTO{
				ID:     t.ID,
				Amount: t.Value,
			},
			FromUser: t.FromUser.Username,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"sent":     sentDTOs,
		"received": receivedDTOs,
	})
}
