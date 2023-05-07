package controllers

import (
	"fmt"
	"net/http"

	"arosara.com/task-manager/db"
	"arosara.com/task-manager/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todoinput models.CreateTodoInput
	if err := c.ShouldBindJSON(&todoinput); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(todoinput)
	todo := models.Todo{Title: todoinput.Title, IsComplete: todoinput.IsComplete}

	err := db.Db.Create(&todo).Error
	if err != nil {
		return
	}
	c.JSON(http.StatusCreated, gin.H{"title": todo.Title, "completed": todo.IsComplete, "id": todo.Id})
}

func GetTodos(c *gin.Context) {

	var todos []models.Todo
	err := db.Db.Find(&todos).Error
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	var todo models.Todo
	if err := db.Db.Where("id=?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
		return
	}
	c.JSON(http.StatusOK, todo)

}
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dbTodo models.Todo
	if err := db.Db.Where("id=?", todo.Id).First(&dbTodo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := db.Db.Save(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todo)

}

func DeleteTodo(c *gin.Context) {

	if err := db.Db.Raw("DELETE FROM todos WHERE id=?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo is deleted successfully"})

}
