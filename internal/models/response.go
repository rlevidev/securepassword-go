package models

type ValidationError struct {
	Errors []string `json:"errors"`
}
