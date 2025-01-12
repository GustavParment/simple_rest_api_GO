package user

import (
	"github.com/gorilla/mux"
	"simplerestapi/api/controllers"
)

func InitGoodbyeRoutes(router *mux.Router) {
	router.HandleFunc("/user/goodbye", controllers.Goodbye).Methods("GET")
}
