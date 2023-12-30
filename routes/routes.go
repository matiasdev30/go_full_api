package routes

import (
	"github.com/gin-gonic/gin"
	docs "github.com/matiasdev30/go_api/docs"
	"github.com/matiasdev30/go_api/handlers"
	"github.com/matiasdev30/go_api/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(g *gin.Engine) {

	basePath := "/go/api/"

	docs.SwaggerInfo.BasePath = basePath

	route := g.Group(basePath)
	{

		auth := route.Group("auth")
		{
			auth.POST("register", handlers.RegisterHanlder)

			auth.POST("login", handlers.LoginHandler)
		}

		task := route.Group("/task", middlewares.Auth())

		{
			task.POST("createTask", handlers.TaskHandler)

			task.GET("getTasks", handlers.GetTaskHanlder)

			task.DELETE("deleteTask", handlers.DeleteTaskHandler)

			task.PUT("updateTask", handlers.UpdateTaskHandler)
		}

	}

	// Initialize swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
