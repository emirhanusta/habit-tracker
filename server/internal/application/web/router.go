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
	habitTracker.Get("/user/:userId", controller.GetUserById)
	habitTracker.Get("/user/email/:email", controller.GetUserByEmail)
	habitTracker.Post("/user", controller.Save)

}
