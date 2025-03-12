package shared

import (
	"encoding/json"
	"net/http"
)

// SetError sets the error message and status code for the JSON response
func (j *JSONRes) SetError(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	j.Err = true
	j.Data = nil
	if err != nil {
		j.Message = err.Error()
	}

	return j.WriteJSON(w, *j, statusCode)
}

// WriteJSON writes a well-structured JSON response to the client.
func (j *JSONRes) WriteJSON(w http.ResponseWriter, payload JSONRes, statusCode int, headers ...http.Header) error {
	w.Header().Set("Content-Type", "application/json")

	// Set custom headers if provided
	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	// Encode payload into JSON
	out, err := json.Marshal(payload)
	if err != nil {
		// Handle JSON encoding errors
		errorResponse := JSONRes{
			Err:     true,
			Message: "Internal Server Error: Failed to encode response",
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return err
	}

	// Set HTTP status code
	w.WriteHeader(statusCode)

	// Write JSON response
	_, err = w.Write(out)
	return err
}
