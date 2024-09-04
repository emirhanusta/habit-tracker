package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"habit-tracker/internal/application/controller/request"
	"habit-tracker/internal/application/controller/response"
	"habit-tracker/internal/application/handler/habit"
	"habit-tracker/internal/application/query"
	"net/http"
)

type IHabitController interface {
	GetAllHabitsByUserId(ctx *fiber.Ctx) error
	GetHabitById(ctx *fiber.Ctx) error
	SaveHabit(ctx *fiber.Ctx) error
	UpdateHabit(ctx *fiber.Ctx) error
	DeleteHabit(ctx *fiber.Ctx) error
}

type habitController struct {
	habitQueryService   query.IHabitQueryService
	habitCommandHAndler habit.ICommandHandler
}

func NewHabitController(habitQueryService query.IHabitQueryService, habitCommandHandler habit.ICommandHandler) IHabitController {
	return &habitController{
		habitQueryService:   habitQueryService,
		habitCommandHAndler: habitCommandHandler,
	}
}

// GetAllHabitsByUserId godoc
// @Summary			This method get all habits by given user id
// @Description		get all habits by given user id
// @Tags			Habit
// @Accept			json
// @Produce			json
// @Param			userId path string true "userId"
// @Success 200 {object} []response.HabitResponse
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/api/v1/habit/user/{userId} [get]
func (h *habitController) GetAllHabitsByUserId(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	if userId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("User id is required")
	}
	fmt.Printf("habitController.GetAllHabitsByUserId: Start getting all habits by userId: %s\n", userId)

	allByUserId, err := h.habitQueryService.GetAllByUserId(ctx.UserContext(), userId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToHabitResponseList(allByUserId))
}

// GetHabitById godoc
// @Summary			This method get habit by given id
// @Description		get habit by id
// @Tags			Habit
// @Accept			json
// @Produce			json
// @Param			id path string true "id"
// @Success 200 {object} response.HabitResponse
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/api/v1/habit/{id} [get]
func (h *habitController) GetHabitById(ctx *fiber.Ctx) error {
	habitId := ctx.Params("id")

	if habitId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("Habit id is required")
	}
	fmt.Printf("habitController.GetHabitById: Start getting byId by id: %s\n", habitId)

	byId, err := h.habitQueryService.GetById(ctx.UserContext(), habitId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToHabitResponse(byId))
}

// SaveHabit godoc
// @Summary			This method save habit
// @Description		save habit
// @Tags			Habit
// @Accept			json
// @Produce			json
// @Param			habit body request.HabitCreateRequest true "habit"
// @Success 200
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/api/v1/habit [post]
func (h *habitController) SaveHabit(ctx *fiber.Ctx) error {
	var habitRequest request.HabitCreateRequest

	if err := ctx.BodyParser(&habitRequest); err != nil {
		fmt.Printf("habitController.SaveHabit: Error parsing request: %v\n", err)
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	fmt.Printf("habitController.SaveHabit: Start saving habit: %v\n", habitRequest)

	if err := h.habitCommandHAndler.Save(ctx.UserContext(), habitRequest.ToCommand()); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON("Habit saved successfully")
}

// UpdateHabit godoc
// @Summary			This method update habit
// @Description		update habit
// @Tags			Habit
// @Accept			json
// @Produce			json
// @Param			requestBody body request.HabitUpdateRequest true "Handle Request Body"
// @Success 200
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/api/v1/habit [put]
func (h *habitController) UpdateHabit(ctx *fiber.Ctx) error {
	var habitRequest request.HabitUpdateRequest

	if err := ctx.BodyParser(&habitRequest); err != nil {
		fmt.Printf("habitController.UpdateHabit: Error parsing request: %v\n", err)
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	fmt.Printf("habitController.UpdateHabit: Start updating habit: %v\n", habitRequest)

	if err := h.habitCommandHAndler.Update(ctx.UserContext(), habitRequest.ToCommand()); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON("Habit updated successfully")
}

// DeleteHabit godoc
// @Summary			This method delete habit by id
// @Description		delete habit by id
// @Tags			Habit
// @Accept			json
// @Produce			json
// @Param			id path string true "id"
// @Success 200
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/api/v1/habit/{id} [delete]
func (h *habitController) DeleteHabit(ctx *fiber.Ctx) error {
	habitId := ctx.Params("id")

	if habitId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("Habit id is required")
	}
	fmt.Printf("habitController.DeleteHabit: Start deleting habit by id: %s\n", habitId)

	if err := h.habitCommandHAndler.Delete(ctx.UserContext(), habitId); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON("Habit deleted successfully")
}
