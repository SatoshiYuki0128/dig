package router

import (
	"dig/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")
	user.GET("/:id", handler.GetOneUserByID)
	user.POST("/", handler.CreateUser)
}
