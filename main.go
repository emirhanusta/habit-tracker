package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"habit-tracker/internal/pkg/postgresql"
	"habit-tracker/internal/pkg/server"
)

func main() {
	app := fiber.New()

	// Connect to PostgreSQL
	_, err := postgresql.ConnectPostgres("postgres", "postgres", "localhost", "5432", "habit-tracker-db")
	if err != nil {
		fmt.Println(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	server.NewServer(app).StartServer()
}
