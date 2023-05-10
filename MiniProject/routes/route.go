package routes

import (
	"MiniProject/constant"
	"MiniProject/controller"
	"MiniProject/midleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo) {

	midleware.LogMiddleware(e)
	// routing with query parameter
	e.POST("/users", controller.CreateUserController)
	e.POST("/users/login", controller.LoginUserController)
	e.POST("/admin/login", controller.LoginAdminController)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.GET("/admin/allusers", controller.GetUsersController)
	eJwt.GET("/admin/:id", controller.GetUserController)
	eJwt.DELETE("/:id", controller.DeleteUserController)
	eJwt.PUT("/admin/:id", controller.UpdateUserController)
	eJwt.GET("/Products", controller.GetProductsController)
	eJwt.POST("/admin/Products", controller.CreateProductController)
	eJwt.GET("/Products/:id", controller.GetProductController)
	eJwt.DELETE("/admin/Products/:id", controller.DeleteProductController)
	eJwt.PUT("/admin/Products/:id", controller.UpdateProductController)
	eJwt.POST("/user/transaksi", controller.CreateTransacsionController)
	eJwt.PUT("/admin/transaksi/:id", controller.UpdateTransacsionController)
	eJwt.GET("/admin/transaksi", controller.GetTransacsionsController)
	eJwt.GET("/admin/transaksi/:id", controller.GetTransacsionController)
	eJwt.GET("/user/transaksi", controller.GetTransacsionController)
}
