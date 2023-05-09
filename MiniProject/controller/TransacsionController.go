package controller

import (
	"MiniProject/database"
	"MiniProject/models"
	"net/http"

	"github.com/golang-jwt/jwt"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/labstack/echo/v4"
)

var Transacsions []models.Transacsion

// get all Transacsions
func GetTransacsionsController(c echo.Context) error {
	err := database.DB.Preload("User").Preload("Product").Find(&Transacsions).Error
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
func GetTransacsionController(c echo.Context) error {
	// Retrieve the JWT token from the request context and extract the role claim
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	//If token admin
	if role == "admin" {
		id := c.Param("id")
		var Transacsion []models.Transacsion
		if err := database.DB.Where("id = ?", id).Preload("User").Preload("Product").Find(&Transacsion).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":   "success get Transaksi by id",
			"Transaksi": Transacsion,
		})
	}

	// If token User
	userId := int(claims["user_id"].(float64))
	var Transacsions []models.Transacsion
	if err := database.DB.Where("user_id = ?", userId).Preload("User").Preload("Product").Find(&Transacsions).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success get user info Transaksi By Id",
		"transactions": Transacsions,
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
	if err := database.DB.Preload("Product").Preload("User").Model(&models.Transacsion{}).Where("id = ?", id).First(&Transacsion).Error; err != nil {
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
	productID := Transacsion.ProductID
	var product models.Product
	if err := database.DB.Model(&models.Product{}).Where("id = ?", productID).First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	// Min -1 Product stock
	product.Stock -= 1

	if err := database.DB.Model(&models.Product{}).Where("id = ?", productID).Update("stock", product.Stock).Error; err != nil {
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
	userId := int(claims["user_id"].(float64))
	// Set the transaction status based on the user role
	if userRole == "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "only user"})
	}
	if userRole == "user" {
		transaction.Status = "pending"
	}

	// Fetch the associated product
	var product models.Product
	if err := database.DB.First(&product, transaction.ProductID).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch product")
	}
	transaction.UserID = userId
	transaction.Amount = product.Price
	transaction.Seller = "Wahyu"

	// Save the transaction to the database
	if err := database.DB.Save(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
	if err := database.DB.Preload("User").First(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user")
	}
	// Prepare the response data
	transacsionResponse := models.TransacsionResponse{
		ID:      int(transaction.ID),
		Product: product.Nama,
		Amount:  transaction.Amount,
		Status:  transaction.Status,
		Seller:  transaction.Seller,
		User:    transaction.User,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Transaction created successfully",
		"transaction": transacsionResponse,
	})
}
