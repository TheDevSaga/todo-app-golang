package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id       int
	Email    string `json:"email"  binding:"required"`
	Name     string `json:"name" binding:"required" `
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Session struct {
	SessionKey string `gorm:"primary_key"`
	UserId     int    `json:"-"`
}

func (user *User) CreateSession(db *gorm.DB) (*Session, error) {
	var err error
	sessenkey := uuid.New().String()
	session := Session{SessionKey: sessenkey, UserId: user.Id}
	err = db.Create(&session).Error

	return &session, err

}
