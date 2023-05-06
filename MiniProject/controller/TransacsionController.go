package controller

import (
	"MiniProject/database"
	"MiniProject/models"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	_ "github.com/jinzhu/gorm/dialects/mysql"

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
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	// Check if the user making the request has the "admin" role
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only admin can perform this action")
	}

	// Update the status to "accept"
	if err := database.DB.Model(&models.Transacsion{}).Where("id = ?", id).Update("status", "accept").Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Transacsion Accept successfully",
		"Transaksi": Transacsion,
	})
}

func CreateTransacsionController(c echo.Context) error {
	// Parse the request body into a transaction object
	var transaction models.Transacsion
	if err := c.Bind(&transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Check the user role from JWT token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userRole := claims["role"].(string)

	// Set the transaction status based on the user role
	if userRole == "user" {
		transaction.Status = "pending"
	}

	// Fetch the associated product
	var product models.Product
	if err := database.DB.First(&product, transaction.ProductID).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch product")
	}
	transaction.Amount = product.Price
	transaction.Seller = "Wahyu"

	// Save the transaction to the database
	if err := database.DB.Save(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	// Prepare the response data
	transacsionResponse := models.TransacsionResponse{
		ID:      int(transaction.ID),
		Product: product.Nama,
		Amount:  transaction.Amount,
		Status:  transaction.Status,
		Seller:  transaction.Seller,
		User:    models.User{},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Transaction created successfully",
		"transaction": transacsionResponse,
	})
}
