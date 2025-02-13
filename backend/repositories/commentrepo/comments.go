package commentrepo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AllPosts returns a slice of all posts
func (p *CommentRepo) ListCommentsByPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	var req CommmentRequest

	posts, err := p.comment.ListCommentsByPost(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = posts
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// CreateComment creates a new comment
func (p *CommentRepo) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	var req Comment

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	post, err := p.comment.CreateComment(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = post
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// AddLike adds a like to a post
func (p *CommentRepo) AddLike(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into a CommentByIDRequest struct
	var req CommentByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		p.res.SetError(w, fmt.Errorf("invalid request payload: %v", err), http.StatusBadRequest)
		return
	}

	// Call the AddLike method to add a like to the post
	err := p.comment.AddLike(&req)
	if err != nil {
		p.res.SetError(w, fmt.Errorf("failed to add like to post: %v", err), http.StatusBadRequest)
		return
	}

	// Prepare and send the success response
	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = nil
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// Dislike removes a like from a post
func (p *CommentRepo) Dislike(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into a CommentByIDRequest struct
	var req CommentByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		p.res.SetError(w, fmt.Errorf("invalid request payload: %v", err), http.StatusBadRequest)
		return
	}

	// Call the AddLike method to add a like to the post
	err := p.comment.DisLike(&req)
	if err != nil {
		p.res.SetError(w, fmt.Errorf("failed to add like to post: %v", err), http.StatusBadRequest)
		return
	}

	// Prepare and send the success response
	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = nil
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}
