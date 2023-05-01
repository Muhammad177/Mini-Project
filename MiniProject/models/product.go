package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	IDProduct   int    `json:"IdProduct" form:"IdProduct"`
	Nama        string `json:"nama" form:"nama"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	Size        int    `json:"size" form:"size"`
	Colour      string `json:"colour" form:"colour"`
	Brand       string `json:"brand" form:"brand"`
	Stock       int    `json:"stock" form:"stock"`
}
