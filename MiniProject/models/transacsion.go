package models

import "github.com/jinzhu/gorm"

type Transacsion struct {
	gorm.Model
	Amount  int     `form:"amount" json:"amount"`
	Status  string  `form:"status" json:"status"`
	Seller  string  `form:"seller" json:"seller"`
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
	User    User    `json:"user"`
}
type TransacsionResponse struct {
	ID      int    `json:"id" form:"name"`
	Product string `form:"product" json:"product"`
	Amount  int    `form:"amount" json:"amount"`
	Status  string `form:"status" json:"status"`
	Seller  string `form:"seller" json:"seller"`
	User    User   `json:"user"`
}
