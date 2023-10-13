package main

import (
	"dig/common"
	"dig/config"
	"dig/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// init config
	err := config.InitConfig()
	if err != nil {
		return
	}

	// init mysql
	err = common.InitMysql()
	if err != nil {
		return
	}

	r := gin.Default()
	v1 := r.Group("/v1")
	router.UserRouter(v1)

	err = r.Run()
	if err != nil {
		return
	}
}
