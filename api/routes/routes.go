package routes

import (
	"github.com/gorilla/mux"
	_ "net/http"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	InitHelloRoutes(router)
	InitGoodbyeRoutes(router)
	InitAuthRoutes(router)

	return router
}
