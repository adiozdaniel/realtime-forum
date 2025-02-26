package postrepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// AllPosts returns a slice of all posts
func (p *PostsRepo) AllPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	posts, err := p.post.ListPosts()
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = posts
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// CreatePost creates a new post
func (p *PostsRepo) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	var req Post

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	post, err := p.post.CreatePost(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = post
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// DeletePost deletes a post
func (p *PostsRepo) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Post

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	err = p.post.DeletePost(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = nil
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// AddLike adds a like to a post
func (p *PostsRepo) PostAddLike(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into a PostLike struct
	var req Like
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		p.res.SetError(w, fmt.Errorf("invalid request payload: %v", err), http.StatusBadRequest)
		return
	}

	// Call the AddLike method to add a like to the post
	post, err := p.post.PostAddLike(&req)
	if err != nil {
		p.res.SetError(w, fmt.Errorf("failed to add like to post: %v", err), http.StatusBadRequest)
		return
	}

	// Prepare and send the success response
	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = post
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// AddLike adds a like to a post
func (p *PostsRepo) CommentAddLike(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into a PostLike struct
	var req Like
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		p.res.SetError(w, fmt.Errorf("invalid request payload: %v", err), http.StatusBadRequest)
		return
	}

	// Call the AddLike method to add a like to the post
	commentLike, err := p.post.CommentAddLike(&req)
	if err != nil {
		p.res.SetError(w, fmt.Errorf("failed to add like to post: %v", err), http.StatusBadRequest)
		return
	}

	// Prepare and send the success response
	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = commentLike
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// Dislike removes a like from a post
func (p *PostsRepo) PostDislike(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into a PostLike struct
	var req Like
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		p.res.SetError(w, fmt.Errorf("invalid request payload: %v", err), http.StatusBadRequest)
		return
	}

	// Call the PostDisLike method to remove dislike a post
	dislike, err := p.post.PostDisLike(&req)
	if err != nil {
		p.res.SetError(w, fmt.Errorf("failed to dislike post: %v", err), http.StatusInternalServerError)
		return
	}

	// Prepare and send the success response
	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = dislike
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// CreateComment creates a new comment
func (p *PostsRepo) CreatePostComment(w http.ResponseWriter, r *http.Request) {
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

	post, err := p.post.CreatePostComment(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = post
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// Posts image uploads a post image and stores on the server
func (p *PostsRepo) UploadPostImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		p.res.SetError(w, errors.New("wrong format"), http.StatusBadRequest)
		return
	}

	postId, _ := p.shared.GenerateUUID()
	image, err := p.shared.SaveImage(r, postId)
	if err != nil {
		p.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	data := map[string]string{
		"img":     image,
		"post_id": postId,
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = data
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// CreateReply creates a new reply
func (p *PostsRepo) CreatePostReply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	var req Reply

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	post, err := p.post.CreateCommentReply(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = post
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// CheckNotifications checks for new notifications
func (p *PostsRepo) CheckNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := p.auth.GetUserIDFromContext(r.Context())
	if !ok {
		p.res.SetError(w, errors.New("not logged in"), http.StatusUnauthorized)
		return
	}

	notifications, err := p.post.GetNotificationsByUserID(userID)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = notifications
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}
