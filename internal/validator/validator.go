package validator

import (
	"regexp"
	"unicode"
)

func ValidatePassword(password string) []string {
	var errors []string

	// Check length
	if len(password) < 8 {
		errors = append(errors, "password must be at least 8 characters long")
	}

	// Check for uppercase
	hasUpper := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		errors = append(errors, "password must contain at least one uppercase letter")
	}

	// Check for lowercase
	hasLower := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLower = true
			break
		}
	}
	if !hasLower {
		errors = append(errors, "password must contain at least one lowercase letter")
	}

	// Check for digit
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		errors = append(errors, "password must contain at least one digit")
	}

	// Check for special character
	specialCharRegex := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`)
	if !specialCharRegex.MatchString(password) {
		errors = append(errors, "password must contain at least one special character")
	}

	return errors
}
