package main

import (
	"AvitoShop/internal/config"
	"AvitoShop/internal/models"
	"AvitoShop/internal/transport/rest"
	"errors"
	"fmt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

func DBinit(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.MoneyTransaction{}, models.BoughtItem{})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	var conf = config.Get()
	log.SetLevel(log.DEBUG)
	log.Debug(conf)

	db_url := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresHost,
		conf.PostgresPort,
		conf.PostgresDatabase)

	log.Debug(db_url)
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DBinit(db)

	e := echo.New()
	h := &handlers.Handler{
		DB:     db,
		Secret: []byte(conf.SecretKey),
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	rApi := e.Group("/api")

	rApi.POST("/auth", h.Auth)
	authRequired := rApi.Group("")
	authRequired.Use(echojwt.JWT(h.Secret))

	authRequired.GET("/info", h.Information)
	authRequired.POST("/sendCoin", h.SendMoney)
	authRequired.GET("/buy/:item", h.BuyItem)
	authRequired.GET("/coinHistory", h.MoneyHistory)

	address := fmt.Sprintf("%s:%s", conf.Host, conf.Port)

	if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
