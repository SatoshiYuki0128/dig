package main

import (
	"dig/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")

	router.UserRouter(v1)

	err := r.Run()
	if err != nil {
		return
	}
}
