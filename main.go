package main

import (
	"fizzBuzz/handlers"
	"fizzBuzz/routes"
	"log"
	"net/http"
)

func init() {
	// Initialise le gestionnaire de statistiques
	handlers.InitStatisticsHandler()

	routes.SetStatisticsHandler(handlers.GetStatisticsHandler())

}
func main() {

	// Configure les routes
	r := routes.ConfigureRoutes()

	// Démarre le serveur HTTP
	log.Fatal(http.ListenAndServe(":8000", r))
}
