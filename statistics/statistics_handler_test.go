package statistics

import (
	"encoding/json"
	"fizzBuzz/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatisticsHandler_GetStatistics(t *testing.T) {
	statsHandler := NewStatisticsHandler()

	request1 := models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"}
	statsHandler.UpdateStatistics(request1)
	request2 := models.FizzBuzzRequest{Int1: 2, Int2: 7, Limit: 10, Str1: "foo", Str2: "bar"}
	statsHandler.UpdateStatistics(request2)

	req, err := http.NewRequest(http.MethodGet, "/statistics", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(statsHandler.GetStatistics)
	handler.ServeHTTP(recorder, req)

	// Vérification du code de réponse
	assert.Equal(t, http.StatusOK, recorder.Code, "Code de réponse incorrect")

	// Vérification du corps de la réponse
	responseBody := recorder.Body.String()
	assert.NotEmpty(t, responseBody, "Corps de réponse vide")

	// Vérification du JSON de réponse
	response := models.StatisticsResponse{
		MostUsedRequest: request1,
		RequestCount: []models.RequestCountPair{
			{Request: request1, Count: 1},
			{Request: request2, Count: 1},
		},
	}
	expectedJSON, err := json.Marshal(response)
	assert.NoError(t, err)

	actualJSON := responseBody
	log.Println("Expected JSON:", string(expectedJSON))
	log.Println("Actual JSON:", actualJSON)

	assert.JSONEq(t, string(expectedJSON), actualJSON, "JSON de réponse incorrect")
}
