package user

import "air-bnb/entities"

type UserFormatResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginUserFormatResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type GetAllUserFormatResponse struct {
	Message string          `json:"message"`
	Data    []entities.User `json:"data"`
}

type GetUserFormatResponse struct {
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}

type RegisterUserFormatResponse struct {
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}

type DeleteUserFormatResponse struct {
	Message string `json:"message"`
}
