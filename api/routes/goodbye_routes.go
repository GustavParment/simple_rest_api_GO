package routes

import (
	"github.com/gorilla/mux"
	"simplerestapi/api/controllers"
)

func InitGoodbyeRoutes(router *mux.Router) {
	router.HandleFunc("/goodbye", controllers.Goodbye).Methods("GET")
}
