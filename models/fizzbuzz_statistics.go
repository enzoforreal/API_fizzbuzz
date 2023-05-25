package models

type StatisticsResponse struct {
	MostUsedRequest FizzBuzzRequest
	RequestCount    []RequestCountPair
}

type RequestCountPair struct {
	Request FizzBuzzRequest
	Count   int
}
