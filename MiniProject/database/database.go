package database

import (
	"fmt"
	"os"

	"MiniProject/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var (
	DB_Username = os.Getenv("DB_Username")
	DB_Password = os.Getenv("DB_Password")
	DB_Port     = os.Getenv("DB_Port")
	DB_Host     = os.Getenv("DB_Host")
	DB_Name     = os.Getenv("DB_Name")
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		DB_Username: "wahyu",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "127.0.0.1",
		DB_Name:     "miniproject",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Transacsion{})
}
