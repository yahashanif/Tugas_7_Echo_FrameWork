package main

import (
	db "pnp/echo-rest/db"
	routes "pnp/echo-rest/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))

}
