package handlers

import (
	"AvitoShop/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MoneyTransactionDTO struct {
	Amount uint `json:"amount"`
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
	username := c.QueryParam("user")

	if username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "username is required"})
	}

	sended, err := services.SendedMoney(h.DB, username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	received, err := services.RecievedMoney(h.DB, username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	sentDTOs := make([]SentMoneyTransactionDTO, 0, len(sended))
	for _, t := range sended {
		sentDTOs = append(sentDTOs, SentMoneyTransactionDTO{
			MoneyTransactionDTO: MoneyTransactionDTO{
				Amount: t.Value,
			},
			ToUser: t.ToUser.Username,
		})
	}

	receivedDTOs := make([]ReceivedMoneyTransactionDTO, 0, len(received))
	for _, t := range received {
		receivedDTOs = append(receivedDTOs, ReceivedMoneyTransactionDTO{
			MoneyTransactionDTO: MoneyTransactionDTO{
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
