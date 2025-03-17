package renders

import "forum/forumapp"

// RendersRepo represents a response implementations
type RendersRepo struct {
	app *forumapp.ForumApp
}

// NewResponse creates a new instance of Response
func NewRendersRepo(app *forumapp.ForumApp) *RendersRepo {
	return &RendersRepo{app}
}
