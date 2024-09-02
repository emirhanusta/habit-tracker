package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"habit-tracker/internal/application/controller/request"
	"habit-tracker/internal/application/controller/response"
	"habit-tracker/internal/application/handler/user"
	"habit-tracker/internal/application/query"
	"net/http"
)

type IUserController interface {
	GetAllUsers(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
	GetUserByEmail(ctx *fiber.Ctx) error
	SaveUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
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

// GetAllUsers godoc
//
//	@Summary		This method get all users
//	@Description	get all users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//
// @Success 200 {object} []response.UserResponse
//
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/api/v1/habit-tracker/user [get]
func (u *userController) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := u.userQueryService.GetAll(ctx.UserContext())

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToResponseList(users))
}

// GetUserById godoc
//
//	@Summary		This method get user by given id
//	@Description	get user by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"userId"
//
// @Success 200 {object} response.UserResponse
//
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/api/v1/habit-tracker/user/{userId} [get]
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

	return ctx.Status(http.StatusOK).JSON(response.ToResponse(byId))
}

// GetUserByEmail godoc
//
//	@Summary		This method get user by given email
//	@Description	get user by email
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			email	path		string	true	"email"
//
// @Success 200 {object} response.UserResponse
//
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/api/v1/habit-tracker/user/email [get]
func (u *userController) GetUserByEmail(ctx *fiber.Ctx) error {
	email := ctx.Query("email")

	if email == "" {
		return ctx.Status(http.StatusBadRequest).JSON("email is required")
	}

	fmt.Printf("userController.GetUserById Started with userId: %s\n", email)

	byEmail, err := u.userQueryService.GetByEmail(ctx.UserContext(), email)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ToResponse(byEmail))
}

// Save godoc
//
//	@Summary		This method used for saving new user
//	@Description	saving new user
//
// @Param requestBody body request.UserCreateRequest nil "Handle Request Body"
//
//	@Tags			User
//	@Accept			json
//	@Produce		json
//
// @Success 200
//
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/api/v1/habit-tracker/user [post]
func (u *userController) SaveUser(ctx *fiber.Ctx) error {
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

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @Tags User
// @Accept json
// @Produce json
// @Param requestBody body request.UserUpdateRequest true "Handle Request Body"
// @Success 201
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/v1/habit-tracker/user [put]
func (u *userController) UpdateUser(ctx *fiber.Ctx) error {
	var req request.UserUpdateRequest

	if err := ctx.BodyParser(&req); err != nil {
		fmt.Printf("userController.Update Error: %v\n", err)
		return err

	}
	fmt.Printf("userController.Update Started with request: %v\n", req)

	if err := u.userCommandHandler.Update(ctx.UserContext(), req.ToCommand()); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON("User updated successfully")
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by id
// @Tags User
// @Accept json
// @Produce json
// @Param userId path string true "userId"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/v1/habit-tracker/user/{userId} [delete]
func (u *userController) DeleteUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	if userId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("userId is required")
	}

	fmt.Printf("userController.DeleteUser Started with userId: %s\n", userId)

	if err := u.userCommandHandler.Delete(ctx.UserContext(), userId); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON("User deleted successfully")
}
