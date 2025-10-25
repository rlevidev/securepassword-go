package models

type ValidationError struct {
	Errors []string `json:"errors"`
}

type ValidationSuccess struct {
	Message string `json:"message"`
}
