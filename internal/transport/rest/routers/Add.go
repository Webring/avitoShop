package routers

import (
	"AvitoShop/internal/transport/rest/handlers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AddRouters(h *handlers.Handler, e *echo.Echo) {
	rApi := e.Group("/api")

	rApi.POST("/auth", h.Auth)
	authRequired := rApi.Group("")
	authRequired.Use(echojwt.JWT(h.Secret))

	authRequired.GET("/profile", h.Information)
	authRequired.POST("/sendCoin", h.SendMoney)
	authRequired.GET("/buy/:item", h.BuyItem)
}
