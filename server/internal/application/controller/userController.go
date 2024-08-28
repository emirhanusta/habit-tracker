package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"habit-tracker/internal/application/controller/request"
	"habit-tracker/internal/application/handler/user"
	"habit-tracker/internal/application/query"
	"net/http"
)

type IUserController interface {
	GetUserById(ctx *fiber.Ctx) error
	GetUserByEmail(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
}

type userController struct {
	userQueryService   query.IUserQueryService
	userCommandHandler user.ICommandHandler
}

func NewUserController(userQueryService query.IUserQueryService, userCommandHandler user.ICommandHandler) IUserController {
	return &userController{
		userQueryService:   userQueryService,
		userCommandHandler: userCommandHandler,
	}
}

func (u *userController) GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	if userId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("userId is required")
	}

	fmt.Printf("userController.GetUserById Started with userId: %s\n", userId)

	byId, err := u.userQueryService.GetById(ctx.UserContext(), userId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(byId)
}

func (u *userController) GetUserByEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")

	if email == "" {
		return ctx.Status(http.StatusBadRequest).JSON("email is required")
	}

	fmt.Printf("userController.GetUserById Started with userId: %s\n", email)

	byEmail, err := u.userQueryService.GetByEmail(ctx.UserContext(), email)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(byEmail)
}

func (u *userController) Save(ctx *fiber.Ctx) error {
	var req request.UserCreateRequest

	if err := ctx.BodyParser(&req); err != nil {
		fmt.Printf("userController.Save Error: %v\n", err)
		return err

	}
	fmt.Printf("userController.Save Started with request: %v\n", req)

	if err := u.userCommandHandler.Save(ctx.UserContext(), req.ToCommand()); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON("User created successfully")
}
