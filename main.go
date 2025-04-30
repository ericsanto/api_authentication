package main

import (
	"github.com/ericsanto/api_authentication/config/database"
	"github.com/ericsanto/api_authentication/internal/routers"
)

func main() {

	database.ConnectDatabase()

	routes := routers.Routers()

	routes.Run(":8080")

}
