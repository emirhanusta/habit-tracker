package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"habit-tracker/internal/application/repository"
	"habit-tracker/internal/domain"
	"habit-tracker/internal/pkg/postgresql"
	"habit-tracker/internal/pkg/server"
)

func main() {
	app := fiber.New()

	// Connect to PostgreSQL
	dbConn, err := postgresql.ConnectPostgres("postgres", "postgres", "localhost", "5432", "habit-tracker-db")
	if err != nil {
		fmt.Println(err)
	}

	// Create a new user repository
	userRepository := repository.NewUserRepository(dbConn)

	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		user, err := userRepository.GetUserByID(id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(user)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		user := new(domain.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		createdUser, err := userRepository.CreateUser(*user)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(createdUser)
	})

	server.NewServer(app).StartServer()
}
