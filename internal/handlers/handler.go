package handlers

import (
	"encoding/json"
	"net/http"

	"securepassword-go/internal/models"
	"securepassword-go/internal/validator"
)

func ValidatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.PasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	errors := validator.ValidatePassword(req.Password)
	if len(errors) > 0 {
		response := models.ValidationError{Errors: errors}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Senha válida - retorna mensagem de sucesso
	response := models.ValidationSuccess{Message: "Senha válida!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
