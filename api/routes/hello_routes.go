package routes

import (
	"github.com/gorilla/mux"
	"simplerestapi/api/controllers"
)

func InitHelloRoutes(router *mux.Router) {
	router.HandleFunc("/hello", controllers.Hello).Methods("GET")
}
