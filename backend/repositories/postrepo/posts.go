package postrepo

import (
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

	// Extract form values
	title := r.FormValue("title")
	content := r.FormValue("content")
	userID := r.FormValue("user_id")

	// Create the post
	post := &Post{
		UserID:      userID,
		PostTitle:   title,
		PostContent: content,
	}

	err := p.post.CreatePost(post)
	if err != nil {
		p.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	p.res.Err = false
	p.res.Message = "Success"
	p.res.Data = post
	p.res.WriteJSON(w, *p.res, http.StatusOK)
}
