package controller

import (
	"MiniProject/database"
	"MiniProject/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type controller struct{}

func initEcho() *echo.Echo {
	database.InitDB()
	database.InitialMigration()

	e := echo.New()

	return e
}

func (c *controller) GetProductsController(ctx echo.Context) error {
	var products []models.Product

	err := database.DB.Find(&products).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var response []map[string]interface{}
	for _, product := range products {
		item := map[string]interface{}{
			"id":    product.ID,
			"brand": product.Brand,
		}
		response = append(response, item)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all Products",
		"Produk":  response,
	})
}

func TestGetAllProducts_Success(t *testing.T) {
	e := initEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctrl := &controller{}
	if assert.NoError(t, ctrl.GetProductsController(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	}
}
