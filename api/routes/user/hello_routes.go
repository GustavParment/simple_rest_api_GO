package user

import (
	"github.com/gorilla/mux"
	"simplerestapi/api/controllers"
)

func InitHelloRoutes(router *mux.Router) {
	router.HandleFunc("/user/hello", controllers.Hello).Methods("GET")
}
