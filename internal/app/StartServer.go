package server

import (
	"AvitoShop/internal/config"
	"AvitoShop/internal/models"
	"AvitoShop/internal/transport/rest/handlers"
	"AvitoShop/internal/transport/rest/routers"
	"errors"
	"fmt"
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
func Start() {
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

	routers.AddRouters(h, e)

	address := fmt.Sprintf("%s:%s", conf.Host, conf.Port)

	if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}

}
