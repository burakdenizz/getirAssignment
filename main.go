package main

import (
	"fmt"
	"getirAssignment/controllers"
	"getirAssignment/database"
	"getirAssignment/routes"
	"net/http"
)

func main() {
	database.Connect()
	KeyValueStore := map[string]string{}
	controllers.KeyValueStore = KeyValueStore
	defer database.Close()
	routes.Setup()
	fmt.Println("Serving at local host 8080")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
