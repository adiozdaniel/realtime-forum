package postrepo

import (
	"forum/forumapp"
	"forum/repositories/shared"
)

// PostRepo represents posts repository
type PostsRepo struct {
	app     *forumapp.ForumApp
	res     *shared.JSONRes
	service *PostService
	shared  *shared.SharedConfig
}

// NewPost returns a new Post instance
func NewPostsRepo(app *forumapp.ForumApp) *PostsRepo {
	postRepo := NewPostRepository(app.Db.Query)
	// Create postService (which depends on postRepo)
	postService := NewPostService(postRepo)

	return &PostsRepo{
		app:     app,
		res:     shared.NewJSONRes(),
		service: postService,
		shared:  shared.NewSharedConfig(),
	}
}
