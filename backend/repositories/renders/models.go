package renders

import "forum/forumapp"

// Response represents a response implementations
type Response struct {
	app *forumapp.ForumApp
}

// NewResponse creates a new instance of Response
func NewResponse(app *forumapp.ForumApp) *Response {
	return &Response{app}
}
