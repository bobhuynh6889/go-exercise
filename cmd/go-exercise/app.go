package main

import (
	db "github.com/bobhuynh6889/go-exercise/internal/app/databases"
	"github.com/bobhuynh6889/go-exercise/internal/app/routers"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379", // Redis server address
	// 	Password: "",               // Redis server password
	// 	DB:       0,                // Default database
	// })
	// defer redisClient.Close()
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

	utils.StartServerWithGracefulShutdown(app)
}
