package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Nama        string `json:"nama" form:"nama"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	Size        int    `json:"size" form:"size"`
	Colour      string `json:"colour" form:"colour"`
	Brand       string `json:"brand" form:"brand"`
	Stock       int    `json:"stock" form:"stock"`
}

type ClassProduct struct {
	ID    int    `json:"id" form:"name"`
	Brand string `json:"brand" form:"brand"`
}
