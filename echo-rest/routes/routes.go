package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Init ...
func Init() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())

	//UserRoute ...
	UserRoute(e.Group("/user"))

	//CustomersRoute ...
	CustomersRoute(e.Group("/customers"))

	//EmployeesRoute ...
	EmployeesRoute(e.Group("/employees"))

	//EmployeesRoute ...
	SuppliersRoute(e.Group("/suppliers"))

	return e
}
