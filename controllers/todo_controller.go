package controllers

import (
	"net/http"

	"arosara.com/task-manager/db"
	"arosara.com/task-manager/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todoinput models.CreateTodoInput
	if err := c.ShouldBindJSON(&todoinput); err != nil {
		c.JSON(http.StatusBadRequest, models.InvalidParameterResponse)
		return
	}
	todo := models.Todo{Title: todoinput.Title, IsComplete: todoinput.IsComplete, UserId: c.GetInt("user_id")}

	err := db.Db.Create(&todo).Error
	if err != nil {
		c.JSON(http.StatusBadGateway, models.BadGatewayResponse)
		return
	}
	c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "Todo is Created"})
}

func GetTodos(c *gin.Context) {

	var todos []models.Todo
	err := db.Db.Where("user_id=?", c.GetInt("user_id")).Find(&todos).Error
	if err != nil {
		c.JSON(http.StatusBadGateway, models.BadGatewayResponse)
		return
	}
	c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "OK", Data: todos})
}

func GetTodoById(c *gin.Context) {
	var todo models.Todo
	if err := db.Db.Where("user_id=?", c.GetInt("user_id")).Where("id=?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, models.BaseResponse{Status: http.StatusNotFound, Message: "No todo found", Error: "No todo found"})
		return
	}
	c.JSON(http.StatusOK, models.BaseResponse{Data: todo, Status: http.StatusOK, Message: "Ok"})

}
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, models.InvalidParameterResponse)
		return
	}
	var dbTodo models.Todo
	if err := db.Db.Where("user_id=?", c.GetInt("user_id")).Where("id=?", todo.Id).First(&dbTodo).Error; err != nil {
		c.JSON(http.StatusNotFound, models.BaseResponse{Status: http.StatusNotFound, Message: "Todo Nod found", Error: "Todo Not found"})
		return
	}
	if err := db.Db.Save(&todo).Error; err != nil {
		c.JSON(http.StatusBadGateway, models.BadGatewayResponse)
	}

	c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "Todo Updated", Data: todo})

}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := db.Db.Where("user_id=?", c.GetInt("user_id")).Where("id=?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, models.BaseResponse{Status: http.StatusNotFound, Message: "Todo Not Found", Error: "Todo Not Found"})
		return
	}
	if err := db.Db.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusBadGateway, models.BadGatewayResponse)
		return
	}

	c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "Todo Deleted Successfully"})

}
