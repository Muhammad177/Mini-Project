package models

import "github.com/jinzhu/gorm"

type Transacsion struct {
	gorm.Model
	UserID    int     `json:"user_id" form:"user_id"`
	ProductID int     `json:"product_id" form:"product_id"`
	Amount    int     `form:"amount" json:"amount"`
	Status    string  `form:"status" json:"status"`
	Seller    string  `form:"seller" json:"seller"`
	Product   Product `json:"product"`
	User      User    `json:"user"`
}

type TransacsionResponse struct {
	ID      int    `json:"id" form:"name"`
	Product string `form:"product" json:"product"`
	Amount  int    `form:"amount" json:"amount"`
	Status  string `form:"status" json:"status"`
	Seller  string `form:"seller" json:"seller"`
	User    User   `json:"user"`
}
