package authrepo

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// RegisterHandler handles user registration.
func (h *AuthRepo) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("hacking detected, method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Extract form values
	var req User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.res.SetError(w, errors.New("hacking detected, invalid request"), http.StatusBadRequest)
		return
	}

	// Register the user
	err = h.user.Register(&req)
	if err != nil {
		if strings.Contains(err.Error(), "oops") {
			h.res.SetError(w, err, http.StatusInternalServerError)
			return
		}

		h.res.SetError(w, err, http.StatusConflict)
		return
	}

	// Generate and store session
	h.auth.GenerateToken(req.UserID, w)

	// Respond with success and token
	h.res.Data = req
	h.res.Err = false
	h.res.Message = "User registered successfully"

	// Respond with JSON
	if err := h.res.WriteJSON(w, *h.res, http.StatusCreated); err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}
}

// LoginHandler handles user login.
func (h *AuthRepo) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("hacking detected, method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.res.SetError(w, errors.New("hacking detected, invalid request"), http.StatusBadRequest)
		return
	}

	user, err := h.user.Login(req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "oops") {
			h.res.SetError(w, err, http.StatusInternalServerError)
			return
		}

		h.res.SetError(w, err, http.StatusUnauthorized)
		return
	}

	// Check if user already has a valid session
	if cookie, err := r.Cookie("session_token"); err == nil {
		if h.auth.ValidateSession(user.UserID, cookie.Value) {
			h.res.Err = false
			h.res.Message = "Login successful (existing session)"
			h.res.Data = user
			h.res.WriteJSON(w, *h.res, http.StatusOK)
			return
		}
	}

	// Generate and store session
	h.auth.GenerateToken(user.UserID, w)
	h.auth.SetUserIDInContext(r.Context(), user.UserID)

	// Respond with success and token
	h.res.Err = false
	h.res.Message = "Login successful"
	h.res.Data = &user

	err = h.res.WriteJSON(w, *h.res, http.StatusOK)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}
}

// LogoutHandler handles user logout.
func (h *AuthRepo) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("hacking detected, method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from request
	userID, ok := h.auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.res.SetError(w, errors.New("not logged in"), http.StatusUnauthorized)
		return
	}

	// Check if user already has a valid session
	if cookie, err := r.Cookie("session_token"); err == nil {
		if h.auth.ValidateSession(userID, cookie.Value) {
			h.auth.Logout(userID)

			h.res.Err = false
			h.res.Message = "Logout successful"
			h.res.Data = nil
			h.res.WriteJSON(w, *h.res, http.StatusOK)
			return
		} else {
			h.res.SetError(w, errors.New("session expired or replaced by a new login"), http.StatusUnauthorized)
			return
		}
	}
}

// CheckAuth confirms if a user is logged in
func (h *AuthRepo) CheckAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.res.SetError(w, errors.New("don't try to hack me"), http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from request
	userID, ok := h.auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.res.SetError(w, errors.New("not logged in"), http.StatusUnauthorized)
		return
	}

	req := User{UserID: userID}
	user, err := h.user.GetUserByID(&req)
	if err != nil {
		h.res.SetError(w, errors.New("oops, something went wrong"), http.StatusInternalServerError)
		return
	}

	// Check if user already has a valid session
	if cookie, err := r.Cookie("session_token"); err == nil {
		if h.auth.ValidateSession(userID, cookie.Value) {
			h.res.Err = false
			h.res.Message = "Logged in"
			h.res.Data = user
			h.res.WriteJSON(w, *h.res, http.StatusOK)
			return
		} else {
			h.res.SetError(w, errors.New("session expired or replaced by a new login"), http.StatusUnauthorized)
			return
		}
	}
}

// UploadProfilePic uploads a profile picture
func (h *AuthRepo) UploadProfilePic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("hacking detected, method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from request
	user_id, ok := h.auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.res.SetError(w, errors.New("not logged in"), http.StatusUnauthorized)
		return
	}

	image, err := h.shared.SaveImage(r, user_id)
	if err != nil {
		h.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	req := User{UserID: user_id, Image: image}

	user, err := h.user.UpdateUser(&req)
	if err != nil {
		if strings.Contains(err.Error(), "oops") {
			h.res.SetError(w, err, http.StatusInternalServerError)
			return
		}

		h.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	h.res.Err = false
	h.res.Message = "Profile picture uploaded successfully"
	h.res.Data = user

	if err := h.res.WriteJSON(w, *h.res, http.StatusOK); err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
	}
}

// UserDashboard returns a user's dashboard
func (h *AuthRepo) UserDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.res.SetError(w, errors.New("hacking detected, method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from request
	userID, ok := h.auth.GetUserIDFromContext(r.Context())
	if !ok {
		h.res.SetError(w, errors.New("not logged in"), http.StatusUnauthorized)
		return
	}

	data, err := h.user.GetUserDashboard(userID)
	if err != nil {
		if strings.Contains(err.Error(), "oops") {
			h.res.SetError(w, err, http.StatusInternalServerError)
			return
		}

		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	h.res.Err = false
	h.res.Message = "Success"
	h.res.Data = data
	h.res.WriteJSON(w, *h.res, http.StatusOK)
}

// EditBio updates a user's bio
func (h *AuthRepo) EditBio(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("hacking detected, method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.res.SetError(w, errors.New("hacking detected, invalid request"), http.StatusBadRequest)
		return
	}

	// Call the EditBio method
	user, err := h.user.EditBio(&req)
	if err != nil {
		if strings.Contains(err.Error(), "oops") {
			h.res.SetError(w, err, http.StatusInternalServerError)
			return
		}

		h.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	// Prepare and send the success response
	h.res.Err = false
	h.res.Message = "Success"
	h.res.Data = user
	h.res.WriteJSON(w, *h.res, http.StatusOK)
}
