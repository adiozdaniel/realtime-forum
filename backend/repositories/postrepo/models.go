package postrepo

import (
	"forum/forumapp"
	"forum/repositories/shared"
	"forum/response"
)

// PostRepo represents posts repository
type PostsRepo struct {
	app    *forumapp.ForumApp
	res    *response.JSONRes
	post   *PostService
	shared *shared.SharedConfig
}

// NewPost returns a new Post instance
func NewPostsRepo(app *forumapp.ForumApp) *PostsRepo {
	postRepo := NewPostRepository(app.Db.Query)
	// Create postService (which depends on postRepo)
	postService := NewPostService(postRepo)

	return &PostsRepo{
		app:    app,
		res:    response.NewJSONRes(),
		post:   postService,
		shared: shared.NewSharedConfig(),
	}
}
