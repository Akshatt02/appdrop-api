package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a standardized error response format for all API errors.
// All error responses follow this structure with error code and human-readable message.
type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// SendError writes a formatted error response to the HTTP response writer.
// Sets appropriate HTTP status code and returns error details in JSON format.
// Example: SendError(w, 404, "NOT_FOUND", "Page not found")
func SendError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var errResp ErrorResponse
	errResp.Error.Code = code
	errResp.Error.Message = message

	json.NewEncoder(w).Encode(errResp)
}

// SendJSON writes a successful response with data to the HTTP response writer.
// Sets Content-Type to application/json and encodes the provided data as JSON.
// Example: SendJSON(w, 200, page) returns the page object with 200 OK status
// Example: SendJSON(w, 201, createdWidget) returns the created widget with 201 Created status
func SendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
