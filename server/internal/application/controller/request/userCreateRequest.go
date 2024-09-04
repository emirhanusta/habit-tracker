package request

import "habit-tracker/internal/application/handler/user"

type UserCreateRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *UserCreateRequest) ToCommand() user.CreateCommand {
	return user.CreateCommand{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
