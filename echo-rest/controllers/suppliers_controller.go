package controllers

import (
	"net/http"
	"pnp/echo-rest/models"

	"github.com/labstack/echo"
)

func FetchAllSuppliers(c echo.Context) (err error) {

	result, err := models.FetchSuppliers()

	return c.JSON(http.StatusOK, result)
}

func AddSuppliers(c echo.Context) (err error) {

	result, err := models.AddSupplier(c)

	return c.JSON(http.StatusOK, result)
}
func UpdateSupplier(c echo.Context) (err error) {

	result, err := models.UpdateSupplier(c)

	return c.JSON(http.StatusOK, result)
}

//DeleteEmployees ...
func DeleteSupplier(c echo.Context) (err error) {

	result, err := models.DeleteSupplier(c)

	return c.JSON(http.StatusOK, result)
}
