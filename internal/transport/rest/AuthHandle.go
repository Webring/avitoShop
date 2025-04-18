package handlers

import (
	"AvitoShop/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"time"

	"net/http"
)

func (h *Handler) Auth(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing username or password")
	}

	var user models.User
	if res := h.DB.First(&user, "username = ?", username); res.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if user.Password != password {
		return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect password")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(h.Secret)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
