package web

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/internal/application/controller"
	"net/http"
)

func InitRouter(app *fiber.App, userController controller.IUserController, habitController controller.IHabitController,
	reminderController controller.IReminderController) {
	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("Server is up and running")
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

	reminder := app.Group("/api/v1/reminder")
	reminder.Get("/habit/:habitId", reminderController.GetAllRemindersByHabitId)
	reminder.Get("/:id", reminderController.GetReminderById)
	reminder.Post("/", reminderController.SaveReminder)
	reminder.Put("/", reminderController.UpdateReminder)
	reminder.Delete("/:id", reminderController.DeleteReminder)
}
