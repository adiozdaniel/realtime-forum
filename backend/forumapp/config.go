package forumapp

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type ForumApp struct {
	Tmpls    *TemplateCache
	Db       *DataConfig
	Sessions sync.Map
	Errors   error
}

func newForumApp() *ForumApp {
	return &ForumApp{
		Tmpls: newTemplateCache(),
		Db:    NewDb(),
	}
}

var (
	instance *ForumApp
	once     sync.Once
)

func ForumInit() (*ForumApp, error) {
	var err error
	once.Do(func() {
		instance = newForumApp()
	})

	err = instance.Tmpls.CreateTemplatesCache()
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (f *ForumApp) GenerateToken(userID string) http.Cookie {
	// Store the token in a session cookie
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    userID,
		Path:     "/",
		HttpOnly: true,             // Prevent JavaScript access
		Secure:   f.IsProduction(), // Secure in production
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	}

	// Store the token in the map
	f.Sessions.Store(userID, cookie)
	return *cookie
}

// generateUUID creates a cryptographically secure random token
func (f *ForumApp) GenerateUUID() (string, error) {
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

// isProduction returns true if the server is running in production mode
func (f *ForumApp) IsProduction() bool {
	return os.Getenv("ENV") == "production"
}
