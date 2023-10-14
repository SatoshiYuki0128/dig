package repository

import (
	"context"
	"dig/internal/common"
	"dig/internal/domain"
	"time"
)

func CreateUser(ctx context.Context, req CreateUserReq) (CreateUserResp, error) {
	user := domain.User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		CreateTime: time.Now(),
	}

	result := common.DB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return CreateUserResp{}, result.Error
	}

	resp := CreateUserResp{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		CreateTime: user.CreateTime,
	}

	return resp, nil
}
