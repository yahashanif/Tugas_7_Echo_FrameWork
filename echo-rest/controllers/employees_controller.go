package controllers

import (
	"net/http"
	"pnp/echo-rest/models"

	"github.com/labstack/echo"
)

//FetchAllCustomers ...
func FetchAllEmployees(c echo.Context) (err error) {

	result, err := models.FetchEmployees()

	return c.JSON(http.StatusOK, result)
}

func AddEmployees(c echo.Context) (err error) {

	result, err := models.AddEmployee(c)

	return c.JSON(http.StatusOK, result)
}
func UpdateEmployee(c echo.Context) (err error) {

	result, err := models.UpdateEmployee(c)

	return c.JSON(http.StatusOK, result)
}

//DeleteEmployees ...
func DeleteEmployee(c echo.Context) (err error) {

	result, err := models.DeleteEmployee(c)

	return c.JSON(http.StatusOK, result)
}
