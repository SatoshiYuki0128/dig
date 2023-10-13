package handler

import (
	"context"
	"dig/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOneUserByID(c *gin.Context) {
	ctx := context.Background()

	id := c.Param("id")

	user, err := usecase.GetOneUserByID(ctx, usecase.GetOneUserByIDReq{ID: id})
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	ctx := context.Background()

	req := usecase.CreateUserReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("c.ShouldBindJSON err: %w", err))
		return
	}

	user, err := usecase.CreateUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
