package response

import "habit-tracker/internal/domain"

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToResponse(user *domain.User) UserResponse {
	return UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func ToResponseList(users []domain.User) []UserResponse {
	var response []UserResponse

	for _, user := range users {
		response = append(response, ToResponse(&user))
	}

	return response
}
