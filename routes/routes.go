package routes

import (
	"getirAssignment/controllers"
	"net/http"
)

func Setup() {
	http.HandleFunc("/api/get-data", controllers.GetData)
	http.HandleFunc("/api/in-memory", controllers.InMemory)
}
