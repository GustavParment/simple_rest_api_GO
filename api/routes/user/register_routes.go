package user

import (
	"github.com/gorilla/mux"
	"simplerestapi/api/controllers"
)

func InitRegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user/register", controllers.Register).Methods("POST")
}
