package shared

import (
	"encoding/json"
	"errors"
	"net/http"
)

// JSONRes represents a JSON response data
type JSONRes struct {
	Err     bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// A constructor function that initializes and returns a new JSONRes instance
func NewJSONRes() *JSONRes {
	return &JSONRes{}
}

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

// ReadJSON intercepts JSON request body to verify validity of the structure
func (j *JSONRes) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("wrong request format")
	}

	maxBytes := 1 * 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return errors.New("wrong request format")
	}

	return nil
}
