package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	db "github.com/bobhuynh6889/go-exercise/internal/app/databases"
	"github.com/bobhuynh6889/go-exercise/internal/app/routers"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New()
	database, err := db.NewPostgreSQLDB()
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}
	defer database.Close()

	routers.InitRoutes(app, database)

	// Start server
	err = app.Listen(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}
	gracefulShutdown(app)
}

func gracefulShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Server is shutting down")

	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server stopped")
}
