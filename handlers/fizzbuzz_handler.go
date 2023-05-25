package handlers

import (
	"encoding/json"
	"fizzBuzz/fizzbuzz"
	"fizzBuzz/models"
	"fizzBuzz/statistics"
	"net/http"
	"strconv"
)

var statisticsHandler *statistics.StatisticsHandler

// InitStatisticsHandler initialise le gestionnaire de statistiques
func InitStatisticsHandler() {
	statisticsHandler = statistics.NewStatisticsHandler()
}

// SetStatisticsHandler définit l'instance du gestionnaire de statistiques pour le gestionnaire FizzBuzzHandler.
func SetStatisticsHandler(sh *statistics.StatisticsHandler) {
	statisticsHandler = sh
}

// GetStatisticsHandler retourne l'instance du gestionnaire de statistiques
func GetStatisticsHandler() *statistics.StatisticsHandler {
	return statisticsHandler
}

// FizzBuzzHandler gère la requête fizz-buzz
// et renvoie la réponse JSON appropriée.
func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	int1Str := r.FormValue("int1")
	int2Str := r.FormValue("int2")
	limitStr := r.FormValue("limit")

	int1, err := strconv.Atoi(int1Str)
	if err != nil || int1 <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	int2, err := strconv.Atoi(int2Str)
	if err != nil || int2 <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	str1 := r.FormValue("str1")
	str2 := r.FormValue("str2")

	request := models.FizzBuzzRequest{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}

	response := fizzbuzz.GenerateFizzBuzz(request)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if statisticsHandler != nil {
		statisticsHandler.UpdateStatistics(request)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
