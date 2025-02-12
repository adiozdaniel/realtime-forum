package postrepo

import (
	"encoding/json"
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

	err = p.post.CreatePost(&req)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = req
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}

// AddLike adds a like to a post
func (p *PostsRepo) AddLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	// Extract form values
	userID := r.FormValue("user_id")
	postID := r.FormValue("post_id")

	post, err := p.post.AddLike(postID, userID)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = post
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}
