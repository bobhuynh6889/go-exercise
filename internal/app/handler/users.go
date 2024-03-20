package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	db "github.com/bobhuynh6889/go-exercise/internal/app/databases"
	"github.com/bobhuynh6889/go-exercise/internal/app/models"
	"github.com/bobhuynh6889/go-exercise/internal/app/queries"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type UserQueries struct {
	DB db.Database
}

func GetListUsers(db db.Database) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		name := c.Query("name")
		userQueries := &queries.UserQueries{DB: db}
		if name == "" {
			users, err := userQueries.GetUsers()
			if err != nil {
				log.Error().Err(err).Msg("Failed to get users")
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			return c.JSON(users)
		} else {
			fmt.Printf("Name: %s\n", name)
			users, err := userQueries.GetUserByName(name)

			if err != nil {
				log.Error().Err(err).Msg("Failed to get users")
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			return c.JSON(users)
		}
	}
}

func GetUserId(db db.Database) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")

		userID, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		userQueries := &queries.UserQueries{DB: db}
		users, err := userQueries.GetUserById(userID)

		if err != nil {
			log.Error().Err(err).Msg("Failed to get users")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(users)
	}
}

func CreateAnUser(db db.Database) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var user models.User

		if err := c.BodyParser(&user); err != nil {
			log.Error().Err(err).Msg("Failed to parse request body")
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
		}
		user.ID = rand.Intn(101)
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		user.CreatedAt = currentTime
		user.UpdatedAt = currentTime

		userQueries := &queries.UserQueries{DB: db}
		if err := userQueries.CreateUser(&user); err != nil {
			log.Error().Err(err).Msg("Failed to create user")
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
		}

		return c.JSON(fiber.Map{"message": "User created successfully", "user": user})
	}
}

func DeleteAnUser(db db.Database) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")

		userID, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		userQueries := &queries.UserQueries{DB: db}
		err = userQueries.DeleteUser(userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Can not delete user"})
		}

		return c.JSON(fiber.Map{"message": "User deleted successfully"})
	}
}

func UpdateAnUser(db db.Database) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")

		userID, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		var updatedUser models.User
		if err := c.BodyParser(&updatedUser); err != nil {
			log.Error().Err(err).Msg("Failed to parse request body")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
		}
		updatedUser.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		var user []models.User
		userQueries := &queries.UserQueries{DB: db}
		user, err = userQueries.UpdateUser(userID, &updatedUser)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Can not update user"})
		}

		return c.JSON(fiber.Map{"message": "User updated successfully", "user": user})
	}
}
