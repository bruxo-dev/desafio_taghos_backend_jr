package utils

import (
	"net/http"
	"regexp"
)

func ValidateInput(input string, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(input)
}

func FormatResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "` + message + `"}`))
}
