package controllers

import (
	"net/http"

	"arosara.com/task-manager/db"
	"arosara.com/task-manager/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request models.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameters"})
		return
	}
	var user models.User
	if err := db.Db.Where("email=?", request.Email).Where("password=?", request.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	session, err := user.CreateSession(db.Db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in creating session key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Successful", "sessionKey": session.SessionKey})

}

func SignUP(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameters"})
		return
	}

	session, err := user.CreateSession(db.Db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in creating session key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SignUp Successful", "sessionKey": session.SessionKey})
}
