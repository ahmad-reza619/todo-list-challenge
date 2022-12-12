package controllers

type FailedResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseActivity[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
