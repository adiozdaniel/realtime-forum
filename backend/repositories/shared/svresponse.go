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

// WriteJSON writes the JSON response to the client
func (j *JSONRes) WriteJSON(w http.ResponseWriter, payload JSONRes, statusCode int, headers ...http.Header) error {
	out, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(out)
	return err
}
