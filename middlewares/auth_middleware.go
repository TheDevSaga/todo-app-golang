package middlewares

import (
	"fmt"
	"net/http"

	"arosara.com/task-manager/db"
	"arosara.com/task-manager/models"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessenkey := c.GetHeader("Authorization")
		fmt.Printf("sessionkey %s", sessenkey)
		if len(sessenkey) < 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "UnAuthorized"})
			return
		}
		var session models.Session
		if err := db.Db.Where("session_key=?", sessenkey).Find(&session).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "UnAuthorized"})
			return
		}
		c.Set("user_id", session.UserId)
		c.Next()

	}
}
