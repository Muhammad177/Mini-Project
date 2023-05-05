package controller

import (
	"MiniProject/database"
	"MiniProject/models"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

var Transacsions []models.Transacsion

// get all Transacsions
func GetTransacsionsController(c echo.Context) error {
	err := database.DB.Preload("User").Find(&Transacsions).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success get all Transacsions",
		"Transaksi": Transacsions,
	})
}

// get Transacsion by id
func GetTransacsionController(c echo.Context) error {
	TransacsionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid Transacsion ID")
	}

	if err := database.DB.Preload("User").First(&Transacsions, TransacsionId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success get Transacsion by id",
		"Transaksi": Transacsions,
	})
}

// delete Transacsion by id
func DeleteTransacsionController(c echo.Context) error {
	id := c.Param("id")
	var Transacsions models.Transacsion
	if err := database.DB.Where("id = ?", id).First(&Transacsions).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Delete(&Transacsions).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success delete Transacsion by id",
		"Transaksi": Transacsions,
	})
}

// update Transacsion by id
func UpdateTransacsionController(c echo.Context) error {
	id := c.Param("id")

	var Transacsion models.Transacsion
	if err := database.DB.Model(&models.Transacsion{}).Where("id = ?", id).First(&Transacsion).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Transacsion not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if err := c.Bind(&Transacsion); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := database.DB.Model(&models.Transacsion{}).Updates(&Transacsion).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Transacsion updated successfully",
		"Transaksi": Transacsion,
	})
}

// create new Transacsion
func CreateTransacsionController(c echo.Context) error {
	Transacsions := models.Transacsion{}
	Products := models.Product{}
	c.Bind(&Transacsions)

	if err := database.DB.Save(&Transacsions).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
TransacsionsResponse := models.TransacsionResponse{
    ID:      int(Transacsions.ID),
    Product: Products.Nama,
    Amount:  Transacsions.Amount,
    Status:  Transacsions.Seller,
    Seller:  "Wahyu",
    User:    models.User{},
}


	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success melakukan transaksi",
		"Transaksi": TransacsionsResponse,
	})
}
