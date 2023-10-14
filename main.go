package main

import (
	"dig/internal/common"
	"dig/internal/config"
	"dig/internal/router"
	"github.com/gin-gonic/gin"
	"log"
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

	r.ForwardedByClientIP = true
	err = r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal("r.SetTrustedProxies error")
		return
	}

	err = r.Run()
	if err != nil {
		log.Fatal("r.Run error")
		return
	}
}
