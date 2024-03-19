package main

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func cacheMiddleware(redisClient *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cacheKey := c.Path()

		cachedData, err := redisClient.Get(c.Context(), cacheKey).Result()
		if err == nil {
			return c.SendString(cachedData)
		}

		err = c.Next()

		if err == nil {
			err = redisClient.Set(c.Context(), cacheKey, c.Response().Body(), time.Minute).Err()
			if err != nil {
				log.Error().Err(err).Msg(err.Error())
			}
		}

		return c.JSON(err)
	}
}
