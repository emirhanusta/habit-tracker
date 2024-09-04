package web

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/internal/application/controller"
	"net/http"
)

func InitRouter(app *fiber.App, userController controller.IUserController, habitController controller.IHabitController) {
	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("OK")
	})

	user := app.Group("/api/v1/user")
	user.Get("/", userController.GetAllUsers)
	user.Get("/email", userController.GetUserByEmail)
	user.Get("/:id", userController.GetUserById)
	user.Post("/", userController.SaveUser)
	user.Put("/", userController.UpdateUser)
	user.Delete("/:id", userController.DeleteUser)

	habit := app.Group("/api/v1/habit")
	habit.Get("/user/:userId", habitController.GetAllHabitsByUserId)
	habit.Get("/:id", habitController.GetHabitById)
	habit.Post("/", habitController.SaveHabit)
	habit.Put("/", habitController.UpdateHabit)
	habit.Delete("/:id", habitController.DeleteHabit)
}
