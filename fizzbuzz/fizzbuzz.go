package fizzbuzz

import (
	"fizzBuzz/models"
	"strconv"
)

//	génère la séquence fizz-buzz
//
// en utilisant les paramètres de requête donnés.
func GenerateFizzBuzz(request models.FizzBuzzRequest) models.FizzBuzzResponse {
	response := models.FizzBuzzResponse{}
	result := make([]string, request.Limit)

	for i := 1; i <= request.Limit; i++ {
		if i%request.Int1 == 0 && i%request.Int2 == 0 {
			result[i-1] = request.Str1 + request.Str2
		} else if i%request.Int1 == 0 {
			result[i-1] = request.Str1
		} else if i%request.Int2 == 0 {
			result[i-1] = request.Str2
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}

	response.Result = result
	return response
}
