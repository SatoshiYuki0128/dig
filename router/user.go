package router

import (
	"dig/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")
	user.GET("/:id", handler.GetOneUserByID)
	user.POST("/", handler.CreateUser)
}
