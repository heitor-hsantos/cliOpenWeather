package routes

import (
	"cliOpn/handlers"
	"github.com/gorilla/mux"
)

// configurar a rota para o link da OpenWheater
func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/Weather", handlers.GetWeatherData).Methods("GET")
	return r
}
