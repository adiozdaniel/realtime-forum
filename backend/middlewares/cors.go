package middlewares

import (
	"net/http"
	"strings"
)

// CorsMiddleware handles CORS headers
func (a *AuthContext) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)

		// Set common CORS headers
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Pass control to the next handler
		next.ServeHTTP(w, r)
	})
}

// AllowedRoutes handles allowed routes
func (a *AuthContext) AllowedRoutes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow static file serving without checking allowed routes
		if strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}

		// Check if the route is allowed
		if ok := a.app.AllowedRoutes[r.URL.Path]; !ok {
			data := map[string]interface{}{
				"Message": r.URL.Path + " is unknown",
			}
			w.WriteHeader(http.StatusNotFound)
			_ = a.render.RenderTemplate(w, "pageNotFound.page.html", data)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

// UserContextMiddleware sets the user ID in the request context if a valid session is found
func (a *AuthContext) UserContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract session token from cookies
		cookie, err := r.Cookie("session_token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		sessionToken := cookie.Value
		var userID string

		// Look for the user ID that matches this session token
		a.Sessions.Range(func(key, value interface{}) bool {
			if value == sessionToken {
				userID = key.(string)
				return false // Stop searching
			}
			return true
		})

		if userID == "" {
			next.ServeHTTP(w, r)
			return
		}

		// Set the user ID in context
		ctx := a.SetUserIDInContext(r.Context(), userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
