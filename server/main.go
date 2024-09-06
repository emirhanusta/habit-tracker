package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"habit-tracker/configuration"
	_ "habit-tracker/docs"
	"habit-tracker/internal/application/controller"
	"habit-tracker/internal/application/handler/habit"
	"habit-tracker/internal/application/handler/reminder"
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
	dbConn, err := postgresql.ConnectPostgres(configuration.PostgresUsername, configuration.PostgresPassword,
		configuration.PostgresqlHost, configuration.PostgresqlPort, configuration.PostgresqlDbName)
	if err != nil {
		fmt.Println(err)
	}

	// Create a user controller
	userRepository := repository.NewUserRepository(dbConn)
	userQueryService := query.NewUserQueryService(userRepository)
	userCommandHandler := user.NewCommandHandler(userRepository)
	userController := controller.NewUserController(userQueryService, userCommandHandler)

	// Create a habit controller
	habitRepository := repository.NewHabitRepository(dbConn)
	habitQueryService := query.NewHabitQueryService(habitRepository)
	habitCommandHandler := habit.NewCommandHandler(habitRepository)
	habitController := controller.NewHabitController(habitQueryService, habitCommandHandler)

	// Create a habit controller
	reminderRepository := repository.NewReminderRepository(dbConn)
	reminderQueryService := query.NewReminderQueryService(reminderRepository)
	reminderCommandHandler := reminder.NewCommandHandler(reminderRepository)
	reminderController := controller.NewReminderController(reminderQueryService, reminderCommandHandler)

	web.InitRouter(app, userController, habitController, reminderController)

	server.NewServer(app).StartServer()
}

func configureSwaggerUi(app *fiber.App) {
	// Swagger injection
	app.Get("/swagger/*", swagger.HandlerDefault)
	// Root path to SwaggerUI redirection
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusMovedPermanently).Redirect("/swagger/index.html")
	})
}
