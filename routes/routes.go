package routes

import (
	"arosara.com/task-manager/controllers"
	"arosara.com/task-manager/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {

	router := gin.Default()
	router.POST("/api/v1/login/", controllers.Login)
	router.POST("/api/v1/sign_up/", controllers.SignUP)

	router.Use(middlewares.AuthMiddleWare())
	router.GET("/api/v1/tasks/", controllers.GetTodos)
	router.POST("/api/v1/tasks/", controllers.CreateTodo)
	router.GET("/api/v1/tasks/:id", controllers.GetTodoById)
	router.PUT("/api/v1/tasks/", controllers.UpdateTodo)
	router.DELETE("/api/v1/tasks/:id", controllers.DeleteTodo)

	router.Run("localhost:3000")
}
