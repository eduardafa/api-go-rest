package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Starting the server with Go")
	routes.HandleRequest()
}
