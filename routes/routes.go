package routes

import (
	"fizzBuzz/handlers"
	"fizzBuzz/statistics"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

var statsHandler *statistics.StatisticsHandler

func SetStatisticsHandler(sh *statistics.StatisticsHandler) {
	statsHandler = sh
}

func ConfigureRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)

	router.HandleFunc("/fizzbuzz", handlers.FizzBuzzHandler).Methods(http.MethodGet)

	router.HandleFunc("/statistics", statsHandler.GetStatistics).Methods(http.MethodGet)

	//Config du middleware CORS
	corsHandler := cors.Default().Handler(router)

	return corsHandler
}
