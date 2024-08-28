package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"habit-tracker/configuration"
	_ "habit-tracker/docs"
	"habit-tracker/internal/application/controller"
	"habit-tracker/internal/application/handler/user"
	"habit-tracker/internal/application/query"
	"habit-tracker/internal/application/repository"
	"habit-tracker/internal/application/web"
	"habit-tracker/internal/pkg/postgresql"
	"habit-tracker/internal/pkg/server"
)

// @title Habit Tracker Fiber REST API
// @version 1.0
// @description This is a sample swagger for Habit Tracker Fiber REST API
// @contact.name emirhan usta
// @contact.email emirhan1usta@gmail.com
func main() {
	app := fiber.New()

	app.Use(recover.New())

	configureSwaggerUi(app)
	// Connect to Postgres
	dbConn, err := postgresql.ConnectPostgres("postgres", "postgres", "localhost", "5432", "habit-tracker-db")
	if err != nil {
		fmt.Println(err)
	}

	// Create a new user repository
	userRepository := repository.NewUserRepository(dbConn)

	// Create a new user query service
	userQueryService := query.NewUserQueryService(userRepository)

	// Create a new user command handler
	userCommandHandler := user.NewCommandHandler(userRepository)

	// Create a new user controller
	userController := controller.NewUserController(userQueryService, userCommandHandler)

	web.InitRouter(app, userController)

	server.NewServer(app).StartServer()
}

func configureSwaggerUi(app *fiber.App) {
	if configuration.Env != "prod" {
		// Swagger injection
		app.Get("/swagger/*", swagger.HandlerDefault)

		// Root path to SwaggerUI redirection
		app.Get("/", func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusMovedPermanently).Redirect("/swagger/index.html")
		})
	}
}
