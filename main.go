package main

import (
	"log"
	"net/http"
	"simplerestapi/api/routes"
)

func main() {

	router := routes.InitRoutes()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
