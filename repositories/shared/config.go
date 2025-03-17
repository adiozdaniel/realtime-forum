package shared

// Shared config represents all the shared logic
type SharedConfig struct {
	json *JSONRes
}

// NewSharedConfig creates a new instance of SharedConfig
func NewSharedConfig() *SharedConfig {
	return &SharedConfig{
		json: NewJSONRes(),
	}
}

// JSONRes represents a JSON response structure.
type JSONRes struct {
	Err     bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// A constructor function that initializes and returns a new JSONRes instance
func NewJSONRes() *JSONRes {
	return &JSONRes{}
}
