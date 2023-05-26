package statistics

import (
	"encoding/json"
	"fizzBuzz/models"
	"net/http"
	"sync"
)

// gestion  des statistiques des requêtes fizz-buzz
type StatisticsHandler struct {
	mu              sync.Mutex
	mostUsedRequest models.FizzBuzzRequest
	requestCounts   []models.RequestCountPair
}

// crée une nouvelle instance de StatisticsHandler
func NewStatisticsHandler() *StatisticsHandler {
	return &StatisticsHandler{
		requestCounts: make([]models.RequestCountPair, 0),
	}
}

func (sh *StatisticsHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	sh.mu.Lock()
	defer sh.mu.Unlock()

	response := models.StatisticsResponse{
		MostUsedRequest: sh.mostUsedRequest,
		RequestCount:    sh.getRequestCountPairs(),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// met à jour les statistiques avec une nouvelle requête fizz-buzz.
func (sh *StatisticsHandler) UpdateStatistics(request models.FizzBuzzRequest) {
	sh.mu.Lock()
	defer sh.mu.Unlock()

	found := false
	for i, pair := range sh.requestCounts {
		if pair.Request == request {
			sh.requestCounts[i].Count++
			found = true
			break
		}
	}

	if !found {
		sh.requestCounts = append(sh.requestCounts, models.RequestCountPair{
			Request: request,
			Count:   1,
		})
	}

	if len(sh.requestCounts) == 1 || sh.requestCounts[len(sh.requestCounts)-1].Count > sh.requestCounts[len(sh.requestCounts)-2].Count {
		sh.mostUsedRequest = request
	}
}

// retourne les paires RequestCount pour la réponse JSON
func (sh *StatisticsHandler) getRequestCountPairs() []models.RequestCountPair {
	pairs := make([]models.RequestCountPair, 0, len(sh.requestCounts))
	for _, pair := range sh.requestCounts {
		pairs = append(pairs, models.RequestCountPair{
			Request: pair.Request,
			Count:   pair.Count,
		})
	}
	return pairs
}
