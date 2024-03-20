package routers

import (
	"context"

	db "github.com/bobhuynh6889/go-exercise/internal/app/databases"
	"github.com/bobhuynh6889/go-exercise/internal/app/handler"
	"github.com/bobhuynh6889/go-exercise/internal/middleware"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type UserQueries struct {
	DB db.Database
}

func InitRoutes(app *fiber.App, db db.Database) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// Test the connection
	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Redis")
	}
	defer redisClient.Close()
	app.Get("/users", middleware.CacheMiddleware(redisClient), handler.GetListUsers(db))
	app.Get("/users/:id", handler.GetUserId(db))
	app.Post("/users", handler.CreateAnUser(db))
	app.Delete("/users/:id", handler.DeleteAnUser(db))
	app.Put("/users/:id", handler.UpdateAnUser(db))
}
