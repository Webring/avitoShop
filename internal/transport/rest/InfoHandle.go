package handlers

import (
	"AvitoShop/internal/models"
	"AvitoShop/internal/services"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Information(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	username, ok := claims["username"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Username not found in token")
	}

	var user models.User

	if res := h.DB.First(&user, "username = ?", username); res.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	inventory, err := services.UserInventory(h.DB, username)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"coins":     user.Money,
		"inventory": inventory,
	})
}
