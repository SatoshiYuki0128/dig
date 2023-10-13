package usecase

import "time"

type GetOneUserByIDReq struct {
	ID string `json:"id"`
}

type CreateUserReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResp struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreateTime time.Time `json:"create_time"`
}
