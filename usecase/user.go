package usecase

import (
	"context"
	"dig/domain"
)

func GetOneUserByID(ctx context.Context, req GetOneUserByIDReq) (domain.User, error) {
	return domain.User{}, nil
}

//TODO get users by filter

func CreateUser(ctx context.Context, req CreateUserReq) (domain.User, error) {
	return domain.User{}, nil
}

//TODO put user

//TODO delete user
