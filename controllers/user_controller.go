package controllers

import (
	"fmt"
	"net/http"

	"arosara.com/task-manager/db"
	"arosara.com/task-manager/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request models.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.InvalidParameterResponse)
		return
	}
	var user models.User
	if err := db.Db.Where("email=?", request.Email).Where("password=?", request.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "Inavlid Credentials", Error: "Invalid Credentials"})
	}

	session, err := user.CreateSession(db.Db)
	if err != nil {
		c.JSON(http.StatusBadGateway, models.BadGatewayResponse)
		return
	}

	c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "Login successfully", Data: session})

}

func SignUP(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, models.InvalidParameterResponse)
		return
	}
	if err := db.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{Status: http.StatusBadRequest, Message: "User Already Exist", Error: "User Already Exist"})
		return
	}

	session, err := user.CreateSession(db.Db)
	if err != nil {
		c.JSON(http.StatusBadGateway, models.BadGatewayResponse)
		return
	}

	c.JSON(http.StatusOK, models.BaseResponse{Status: http.StatusOK, Message: "SignUp Successfully", Data: session})
}
