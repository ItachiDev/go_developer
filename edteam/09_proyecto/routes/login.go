package routes

import (
	"github.com/go_developer/edteam/09_proyecto/controllers"
	"github.com/gorilla/mux"
)

// SetLoginRouter router para el login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
