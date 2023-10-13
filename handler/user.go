package handler

import (
	"context"
	"dig/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/go-playground/validator/v10"
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

	// bind request
	req := CreateUserReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("bind json err: %s", err))
		return
	}

	// validate req
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("validate req err: %s", err))
		return
	}

	user, err := usecase.CreateUser(ctx, req.ToUseCase())
	if err != nil {
		c.JSON(http.StatusOK, fmt.Sprintf("usecase.CreateUser err: %s", err))
		return
	}

	c.JSON(http.StatusOK, user)
}
