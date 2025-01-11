package routes

import (
	"github.com/gorilla/mux"
	"simplerestapi/api/controllers"
)

func InitAuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}
