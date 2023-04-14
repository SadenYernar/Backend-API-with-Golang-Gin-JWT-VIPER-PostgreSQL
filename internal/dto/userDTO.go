package dto

import "project/internal/model"

type UserResponseDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func UserDTO(user model.User) *UserResponseDTO {
	return &UserResponseDTO{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
}
