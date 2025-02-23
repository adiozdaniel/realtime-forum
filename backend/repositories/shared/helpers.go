package shared

import (
	"crypto/rand"
	"fmt"
)

// GenerateUUID creates a cryptographically secure random token
func (s *SharedConfig) GenerateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Format it as a UUID-like string
	token := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])

	return token, nil
}
