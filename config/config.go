package config

import (
	"github.com/gin-gonic/gin"
	"github.com/matiasdev30/go_api/routes"
)

func InitGin(){

	gin := gin.Default()

	routes.Routes(gin)

	gin.Run()
}

