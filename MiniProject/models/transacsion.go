package models

import "github.com/jinzhu/gorm"

type Transacsion struct {
	gorm.Model
	TransacsionID int    `json:"Transacsion_id" form:"Transacsion_id"`
	Amount        string `form:"amount" json:"amount"`
	Status        string `form:"status" json:"status"`
	Seller        string `form:"seller" json:"seller"`
	User          User   `json:"user"`
}
