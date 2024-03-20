package middleware

import (
	"context"
	"fmt"
	"time"

	db "github.com/bobhuynh6889/go-exercise/internal/app/databases"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func CacheMiddleware(redisClient *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()

		cachedData, err := redisClient.Get(ctx, "cached_data").Result()
		if err == redis.Nil {
			data, err := db.NewPostgreSQLDB()
			if err != nil {
				log.Error().Err(err).Msg("Failed to fetch data from database")
				return c.Next()
			}

			err = redisClient.Set(ctx, "cached_data", data, 24*time.Hour).Err()
			if err != nil {
				log.Error().Err(err).Msg("Failed to cache data")
			}
		} else if err != nil {
			log.Error().Err(err).Msg("Error checking cache")
		} else {
			fmt.Println("Data from cache:", cachedData)
			c.Locals("cached_data", cachedData)
		}

		return c.Next()
	}
}
