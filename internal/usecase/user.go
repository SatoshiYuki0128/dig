package usecase

import (
	"context"
	"dig/internal/domain"
	repository2 "dig/internal/repository"
)

func GetOneUserByID(ctx context.Context, req GetOneUserByIDReq) (domain.User, error) {
	return domain.User{}, nil
}

//TODO get users by filter

func CreateUser(ctx context.Context, req CreateUserReq) (CreateUserResp, error) {
	user := repository2.CreateUserReq{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := repository2.CreateUser(ctx, user)
	if err != nil {
		return CreateUserResp{}, err
	}

	resp := CreateUserResp{
		ID:         result.ID,
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		CreateTime: result.CreateTime,
	}

	return resp, nil
}

//TODO put user

//TODO delete user
