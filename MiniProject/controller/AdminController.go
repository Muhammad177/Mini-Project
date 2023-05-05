package controller

import (
	"MiniProject/database"
	"MiniProject/midleware"

	"net/http"

	"MiniProject/models"

	"github.com/golang-jwt/jwt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

var users []models.User

// get all users
func GetUsersController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "only admin can access"})
	}
	err := database.DB.Find(&users).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// Retrieve the JWT token from the request context and extract the role claim
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	//If token admin
	if role == "admin" {
		id := c.Param("id")
		var users models.User
		if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get user by id",
			"user":    users,
		})
	}

	// If token User
	username := claims["name"].(string)
	var users []models.User
	if err := database.DB.Where("name = ?", username).Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user info by name",
		"user":    users,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "only admin can access"})
	}
	id := c.Param("id")
	var users models.User

	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Delete(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
		"user":    users,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// Retrieve the JWT token from the request context and extract the role claim
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if role == "admin" {
		// If token has admin role, update user by ID
		id := c.Param("id")

		var user models.User
		if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
		}

		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
		}

		if err := database.DB.Model(&user).Updates(user).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User updated successfully",
			"user":    user,
		})
	} else {
		// If token doesn't have admin role, update user by token claims name
		username := claims["name"].(string)

		var user models.User
		if err := database.DB.Where("name = ?", username).First(&user).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
		}

		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
		}

		if err := database.DB.Model(&user).Updates(user).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User updated successfully",
			"user":    user,
		})
	}
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := database.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	if err := database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}

	token, err := midleware.CreateToken(int(user.ID), user.Name, "user")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed Login",
			"error":   err.Error(),
		})
	}
	usersResponse := models.UserResponse{int(user.ID), user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    usersResponse,
	})

}
func LoginAdminController(c echo.Context) error {
	admin := models.Admin{ID: 1, Name: "Wahyu", Email: "admin@gmail.com", Password: "admin123"}
	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}
	var admins = []models.Admin{
		{ID: 1, Name: "Wahyu", Email: "admin@gmail.com", Password: "admin123"},
	}
	for _, a := range admins {
		if a.Email == admin.Email && a.Password == admin.Password {
			token, err := midleware.CreateToken(int(a.ID), a.Name, "admin")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Failed Login",
					"error":   err.Error(),
				})
			}

			adminResponse := models.AdminResponse{admin.ID, admin.Name, admin.Email, token}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Login Admin Sukses",
				"Admin":   adminResponse,
			})
		}
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid username or password"})
}
