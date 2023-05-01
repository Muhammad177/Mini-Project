package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Phonenumber int    `json:"phonenumber" form:"phonenumber"`
	Address     string `json:"address" form:"address"`
}
type UserResponse struct {
	ID    int    `json:"id" form:"name"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
