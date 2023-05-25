package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzHandler_Success(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
	assert.NoError(t, err)

	params := url.Values{
		"int1":  []string{"3"},
		"int2":  []string{"5"},
		"limit": []string{"15"},
		"str1":  []string{"fizz"},
		"str2":  []string{"buzz"},
	}
	req.URL.RawQuery = params.Encode()

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzBuzzHandler)
	handler.ServeHTTP(recorder, req)

	expectedJSON := `{"result":["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]}`
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedJSON, recorder.Body.String())
}

func TestFizzBuzzHandler_MissingParameters(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
	assert.NoError(t, err)

	params := url.Values{
		"int1":  []string{"3"},
		"limit": []string{"15"},
		"str1":  []string{"fizz"},
		"str2":  []string{"buzz"},
	}
	req.URL.RawQuery = params.Encode()

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzBuzzHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFizzBuzzHandler_InvalidParameterValues(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
	assert.NoError(t, err)

	params := url.Values{
		"int1":  []string{"3"},
		"int2":  []string{"5"},
		"limit": []string{"0"},
		"str1":  []string{"fizz"},
		"str2":  []string{"buzz"},
	}
	req.URL.RawQuery = params.Encode()

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzBuzzHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFizzBuzzHandler_EmptyParameterValues(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
	assert.NoError(t, err)

	params := url.Values{
		"int1":  []string{""},
		"int2":  []string{"5"},
		"limit": []string{"15"},
		"str1":  []string{"fizz"},
		"str2":  []string{"buzz"},
	}
	req.URL.RawQuery = params.Encode()

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzBuzzHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFizzBuzzHandler_NonNumericParameterValues(t *testing.T) {
	// Préparation de la requête avec les paramètres invalides
	req, err := http.NewRequest(http.MethodGet, "/fizzbuzz?int1=3&int2=invalid&limit=15&str1=fizz&str2=buzz", nil)
	assert.NoError(t, err)

	// Exécution de la requête avec le gestionnaire FizzBuzzHandler
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzBuzzHandler)
	handler.ServeHTTP(recorder, req)

	// Vérification du code de réponse
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFizzBuzzHandler_SpecialCharactersInParameterValues(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
	assert.NoError(t, err)

	params := url.Values{
		"int1":  []string{"3"},
		"int2":  []string{"5"},
		"limit": []string{"15"},
		"str1":  []string{"f!zz"},
		"str2":  []string{"buz@z"},
		"param": []string{"!@#$%^&"},
	}
	req.URL.RawQuery = params.Encode()

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzBuzzHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedJSON := `{"result":["1","2","f!zz","4","buz@z","f!zz","7","8","f!zz","buz@z","11","f!zz","13","14","f!zzbuz@z"]}`
	assert.Equal(t, expectedJSON, recorder.Body.String())
}
