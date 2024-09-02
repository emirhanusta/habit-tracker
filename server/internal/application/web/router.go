package web

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/internal/application/controller"
	"net/http"
)

func InitRouter(app *fiber.App, controller controller.IUserController) {
	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("OK")
	})

	habitTracker := app.Group("/api/v1/habit-tracker")
	habitTracker.Get("/user", controller.GetAllUsers)
	habitTracker.Get("/user/email", controller.GetUserByEmail)
	habitTracker.Get("/user/:userId", controller.GetUserById)
	habitTracker.Post("/user", controller.SaveUser)
	habitTracker.Put("/user", controller.UpdateUser)
	habitTracker.Delete("/user/:userId", controller.DeleteUser)
}
