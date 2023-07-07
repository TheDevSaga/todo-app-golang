package routes

import (
	"arosara.com/task-manager/controllers"
	"arosara.com/task-manager/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {

	router := gin.Default()
	api := router.Group("api")
	{
		auth := api.Group("v1/auth")
		{
			auth.POST("/login/", controllers.Login)
			auth.POST("/sign_up/", controllers.SignUP)
		}

		api.Use(middlewares.AuthMiddleWare())
		task := api.Group("/v1/tasks")
		{
			task.GET("/", controllers.GetTodos)
			task.POST("/", controllers.CreateTodo)
			task.GET("/:id", controllers.GetTodoById)
			task.PUT("/", controllers.UpdateTodo)
			task.DELETE("/:id", controllers.DeleteTodo)
		}

	}

	router.Run("localhost:3000")
}
