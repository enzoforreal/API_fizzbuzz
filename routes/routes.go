package routes

import (
	"fizzBuzz/handlers"
	"fizzBuzz/statistics"
	"net/http"

	"github.com/gorilla/mux"
)

var statsHandler *statistics.StatisticsHandler

// SetStatisticsHandler définit l'instance du gestionnaire de statistiques
func SetStatisticsHandler(sh *statistics.StatisticsHandler) {
	statsHandler = sh
}

func ConfigureRoutes() *mux.Router {
	router := mux.NewRouter()

	// Route Home
	router.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)
	//URL pour le fizz-buzz
	router.HandleFunc("/fizzbuzz", handlers.FizzBuzzHandler).Methods(http.MethodGet)
	//URL des stats
	router.HandleFunc("/statistics", statsHandler.GetStatistics).Methods(http.MethodGet)

	return router
}
