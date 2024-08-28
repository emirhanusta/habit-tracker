package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"habit-tracker/internal/application/controller"
	"habit-tracker/internal/application/handler/user"
	"habit-tracker/internal/application/query"
	"habit-tracker/internal/application/repository"
	"habit-tracker/internal/application/web"
	"habit-tracker/internal/pkg/postgresql"
	"habit-tracker/internal/pkg/server"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
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
