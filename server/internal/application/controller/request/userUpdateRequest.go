package request

import "habit-tracker/internal/application/handler/user"

type UserUpdateRequest struct {
	Id       string `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *UserUpdateRequest) ToCommand() user.Command {
	return user.Command{
		Id:       u.Id,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
