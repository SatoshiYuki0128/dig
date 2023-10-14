package handler

import (
	"dig/internal/usecase"
)

type CreateUserReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *CreateUserReq) ToUseCase() usecase.CreateUserReq {
	return usecase.CreateUserReq{
		Name:     c.Name,
		Email:    c.Email,
		Password: c.Password,
	}
}
