package auth

import (
	"github.com/gorilla/mux"
	"simplerestapi/api/controllers"
)

func InitAuthRoutes(router *mux.Router) {
	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")
}
