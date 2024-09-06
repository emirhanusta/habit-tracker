package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"habit-tracker/internal/application/controller/request"
	"habit-tracker/internal/application/controller/response"
	"habit-tracker/internal/application/handler/reminder"
	"habit-tracker/internal/application/query"
	"net/http"
)

type IReminderController interface {
	GetAllRemindersByHabitId(ctx *fiber.Ctx) error
	GetReminderById(ctx *fiber.Ctx) error
	SaveReminder(ctx *fiber.Ctx) error
	UpdateReminder(ctx *fiber.Ctx) error
	DeleteReminder(ctx *fiber.Ctx) error
}

type reminderController struct {
	reminderQueryService   query.IReminderQueryService
	reminderCommandHandler reminder.ICommandHandler
}

func NewReminderController(reminderQueryService query.IReminderQueryService, reminderCommandHandler reminder.ICommandHandler) IReminderController {
	return &reminderController{
		reminderQueryService:   reminderQueryService,
		reminderCommandHandler: reminderCommandHandler,
	}
}

// GetAllRemindersByHabitId godoc
// @Summary			This method get all reminders by given habit id
// @Description		get all reminders by given habit id
// @Tags			Reminder
// @Accept			json
// @Produce			json
// @Param			habitId path string true "habitId"
// @Success 200 {object} []response.ReminderResponse
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/api/v1/reminder/habit/{habitId} [get]
func (r *reminderController) GetAllRemindersByHabitId(ctx *fiber.Ctx) error {
	habitId := ctx.Params("habitId")

	if habitId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("Habit id is required")
	}
	fmt.Printf("reminderController.GetAllRemindersByHabitId: Start getting all reminders by habitId: %s\n", habitId)

	reminders, err := r.reminderQueryService.GetAllByHabitId(ctx.UserContext(), habitId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToReminderResponseList(reminders))
}

// GetReminderById godoc
// @Summary			This method get reminder by given id
// @Description		get reminder by id
// @Tags			Reminder
// @Accept			json
// @Produce			json
// @Param			id path string true "id"
// @Success 200 {object} response.ReminderResponse
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/api/v1/reminder/{id} [get]
func (r *reminderController) GetReminderById(ctx *fiber.Ctx) error {
	reminderId := ctx.Params("id")

	if reminderId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("Reminder id is required")
	}
	fmt.Printf("reminderController.GetReminderById: Start getting byId by id: %s\n", reminderId)

	byId, err := r.reminderQueryService.GetById(ctx.UserContext(), reminderId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToReminderResponse(byId))
}

// SaveReminder godoc
// @Summary			This method save reminder
// @Description		save reminder
// @Tags			Reminder
// @Accept			json
// @Produce			json
// @Param			reminder body request.ReminderCreateRequest true "reminder"
// @Success 200
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/api/v1/reminder [post]
func (r *reminderController) SaveReminder(ctx *fiber.Ctx) error {
	var reminderRequest request.ReminderCreateRequest

	if err := ctx.BodyParser(&reminderRequest); err != nil {
		fmt.Printf("reminderController.SaveReminder: Error parsing request: %v\n", err)
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	fmt.Printf("reminderController.SaveReminder: Start saving reminder: %v\n", reminderRequest)

	if err := r.reminderCommandHandler.Save(ctx.UserContext(), reminderRequest.ToCommand()); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON("Reminder saved successfully")
}

// UpdateReminder godoc
// @Summary			This method update reminder
// @Description		update reminder
// @Tags			Reminder
// @Accept			json
// @Produce			json
// @Param			requestBody body request.ReminderUpdateRequest true "Handle Request Body"
// @Success 200
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/api/v1/reminder [put]
func (r *reminderController) UpdateReminder(ctx *fiber.Ctx) error {
	var reminderRequest request.ReminderUpdateRequest

	if err := ctx.BodyParser(&reminderRequest); err != nil {
		fmt.Printf("reminderController.UpdateReminder: Error parsing request: %v\n", err)
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	fmt.Printf("reminderController.UpdateReminder: Start updating reminder: %v\n", reminderRequest)

	if err := r.reminderCommandHandler.Update(ctx.UserContext(), reminderRequest.ToCommand()); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON("Reminder updated successfully")
}

// DeleteReminder godoc
// @Summary			This method delete reminder by id
// @Description		delete reminder by id
// @Tags			Reminder
// @Accept			json
// @Produce			json
// @Param			id path string true "id"
// @Success 200
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/api/v1/reminder/{id} [delete]
func (r *reminderController) DeleteReminder(ctx *fiber.Ctx) error {
	reminderId := ctx.Params("id")

	if reminderId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("Reminder id is required")
	}
	fmt.Printf("reminderController.DeleteReminder: Start deleting reminder by id: %s\n", reminderId)

	if err := r.reminderCommandHandler.Delete(ctx.UserContext(), reminderId); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON("Reminder deleted successfully")
}
