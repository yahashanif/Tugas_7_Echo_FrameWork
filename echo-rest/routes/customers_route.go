package routes

import (
	"pnp/echo-rest/controllers"

	"github.com/labstack/echo"
)

//CustomersRoute ...
func CustomersRoute(g *echo.Group) {

	g.POST("/list", controllers.FetchAllCustomers)

	g.POST("/add", controllers.StoreCustomer)

	g.POST("/update", controllers.UpdateCustomer)

	g.POST("/delete", controllers.DeleteCustomer)

}
