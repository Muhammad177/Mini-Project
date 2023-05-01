package routes

import (
	"MiniProject/constant"
	"MiniProject/controller"
	"MiniProject/midleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	midleware.LogMiddleware(e)
	// routing with query parameter
	e.POST("/users", controller.CreateUserController)
	e.POST("/users/login", controller.LoginUserController)
	e.POST("/admin/login", controller.LoginAdminController)
	// routing with query parameter
	e.GET("/Transacsions", controller.GetTransacsionsController)
	e.POST("/Transacsions", controller.CreateTransacsionController)
	e.GET("/Transacsions/:id", controller.GetTransacsionController)
	e.DELETE("/Transacsions/:id", controller.DeleteTransacsionController)
	e.PUT("/Transacsions/:id", controller.UpdateTransacsionController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(middleware.BasicAuth(midleware.BasicAuthDb))
	eAuthBasic.GET("/users", controller.GetUsersController)
	eAuthBasic.GET("/users/:id", controller.GetUserController)
	eAuthBasic.DELETE("/users/:id", controller.DeleteUserController)
	eAuthBasic.PUT("/users/:id", controller.UpdateUserController)
	eAuthBasic.GET("/Products", controller.GetProductsController)
	eAuthBasic.POST("/Products", controller.CreateProductController)
	eAuthBasic.GET("/Products/:id", controller.GetProductController)
	eAuthBasic.DELETE("/Products/:id", controller.DeleteProductController)
	eAuthBasic.PUT("/Products/:id", controller.UpdateProductController)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.GET("/admin", controller.GetUsersController)
	eJwt.GET("/admin/:id", controller.GetUserController)
	eJwt.DELETE("/admin/:id", controller.DeleteUserController)
	eJwt.PUT("/admin/:id", controller.UpdateUserController)
	eJwt.GET("/Products", controller.GetProductsController)
	eJwt.POST("/admin/Products", controller.CreateProductController)
	eJwt.GET("/Products/:id", controller.GetProductController)
	eJwt.DELETE("/Products/:id", controller.DeleteProductController)
	eJwt.PUT("/admin/Products/:id", controller.UpdateProductController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
	return e
}
