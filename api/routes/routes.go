package routes

import (
	"github.com/gorilla/mux"
	_ "net/http"
	"simplerestapi/api/routes/auth"
	"simplerestapi/api/routes/user"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	user.InitHelloRoutes(router)
	user.InitGoodbyeRoutes(router)
	user.InitRegisterRoutes(router)
	auth.InitAuthRoutes(router)

	return router
}
