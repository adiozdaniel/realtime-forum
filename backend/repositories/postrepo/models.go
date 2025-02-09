package postrepo

import (
	"forum/forumapp"
	"forum/repositories/shared"
)

// PostRepo represents posts repository
type PostsRepo struct {
	app    *forumapp.ForumApp
	res    *shared.JSONRes
	post   *PostService
	shared *shared.SharedConfig
}

// NewPost returns a new Post instance
func NewPostsRepo(app *forumapp.ForumApp) *PostsRepo {
	res := shared.NewJSONRes()
	shared := shared.NewSharedConfig()
	postRepo := NewPostRepository(app.Db.Query)
	postService := NewPostService(postRepo)

	return &PostsRepo{
		app:    app,
		res:    res,
		post:   postService,
		shared: shared,
	}
}
