package main

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis server password
		DB:       0,                // Default database
	})
	defer redisClient.Close()
	app := fiber.New()
	// app.Use(cacheMiddleware(redisClient))

	// Database connection
	db, err := sql.Open("postgres", "user=bobhuynh password=postgres dbname=postgres host=192.168.215.2 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}
	defer db.Close()

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	// get
	app.Get("/users", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id, user_name, pass_word FROM users")
		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer rows.Close()

		// declare user Model
		users := make([]User, 0)

		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.UserName, &user.PassWord); err != nil {
				log.Error().Err(err).Msg(err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			users = append(users, user)
		}
		if err := rows.Err(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(users)
	})

	// post
	app.Post("/users", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		insertQuery := `
			INSERT INTO users (id, user_name, pass_word)
			VALUES ($1, $2, $3)
		`
		rand.Seed(time.Now().UnixNano())

		randomInt := rand.Intn(101)

		user.ID = randomInt
		_, err := db.Exec(insertQuery, user.ID, user.UserName, user.PassWord)
		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			return err
		}

		return c.JSON(user)
	})

	// delete
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		userID := c.Params("id")

		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1", userID).Scan(&count)
		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		if count == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}

		deleteStatement := `
        DELETE FROM users
        WHERE id = $1
    `

		_, err = db.Exec(deleteStatement, userID)
		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"message": "User deleted successfully"})
	})

	// patch
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		userID := c.Params("id")

		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1", userID).Scan(&count)
		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		if count == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}

		var updatedUser User
		if err := c.BodyParser(&updatedUser); err != nil {
			log.Error().Err(err).Msg(err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		updateStatement := `
        UPDATE users
        SET user_name = $2, pass_word = $3
        WHERE id = $1
    `

		_, err = db.Exec(updateStatement, userID, updatedUser.UserName, updatedUser.PassWord)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"message": "User updated successfully", "user": updatedUser})
	})

	// Start server
	err = app.Listen(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}
}
