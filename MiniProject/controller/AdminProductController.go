package controller

import (
	"MiniProject/database"

	"MiniProject/models"
	"net/http"

	"github.com/golang-jwt/jwt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

var Products []models.Product

// get all Products
func GetProductsController(c echo.Context) error {

	err := database.DB.Find(&Products).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var classes []models.ClassProduct
	for _, prod := range Products {
		class := models.ClassProduct{ID: int(prod.ID), Brand: prod.Brand}
		classes = append(classes, class)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all Products",
		"Produk":   classes,
	})
}

// get Product by id
func GetProductController(c echo.Context) error {
	id := c.Param("id")
	var Products models.Product
	if err := database.DB.Where("id = ?", id).First(&Products).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Product by id",
		"user":    Products,
	})
}

// delete Product by id
func DeleteProductController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "only admin can access"})
	}
	id := c.Param("id")
	var Products models.Product
	if err := database.DB.Where("id = ?", id).First(&Products).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Delete(&Products).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Product by id",
		"Produk":  Products,
	})
}

// update Product by id
func UpdateProductController(c echo.Context) error {
	id := c.Param("id")

	var product models.Product
	if err := database.DB.Model(&models.Product{}).Where("id = ?", id).First(&product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := database.DB.Model(&models.Product{}).Updates(product).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product updated successfully",
		"user":    product,
	})
}

// create new Product
func CreateProductController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "only admin can access"})
	}
	Products := models.Product{}
	c.Bind(&Products)

	if err := database.DB.Save(&Products).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new Product",
		"Produk":  Products,
	})
}
