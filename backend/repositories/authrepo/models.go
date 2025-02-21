package authrepo

import (
	"forum/forumapp"
	"forum/middlewares"
	"forum/repositories/postrepo"
	"forum/repositories/shared"
	"time"
)

// LoginRequest represents the request body for user login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LogoutRequest represents the request body for logout
type LogoutRequest struct {
	UserId string `json:"user_id"`
}

// User represents a user in the database
type User struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UserName  string    `json:"user_name"`
	Image     string    `json:"image"`
	Role      string    `json:"role"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserData represents the user data
type UserData struct {
	UserInfo       *User               `json:"user_info"`
	Posts          []*postrepo.Post    `json:"posts"`
	Comments       []*postrepo.Comment `json:"comments"`
	Replies        []*postrepo.Reply   `json:"replies"`
	LikedPosts     []*postrepo.Post    `json:"liked_posts"`
	LikedComments  []*postrepo.Comment `json:"liked_comments"`
	Likes          []*postrepo.Like    `json:"likes"`
	Dislikes       []*postrepo.Like    `json:"dislikes"`
	RecentActivity []*RecentActivity   `json:"recent_activity"`
}

// RecentActivity represents the recent activity
type RecentActivity struct {
	Post    postrepo.Post    `json:"post"`
	Comment postrepo.Comment `json:"comment"`
	Reply   postrepo.Reply   `json:"reply"`
	Like    postrepo.Like    `json:"like"`
	Dislike postrepo.Like    `json:"dislike"`
}

// AuthRepo represents the repository for authentication
type AuthRepo struct {
	app    *forumapp.ForumApp
	res    *shared.JSONRes
	user   *UserService
	shared *shared.SharedConfig
	auth   *middlewares.AuthContext
}

// NewAuthRepo creates a new instance of AuthRepo
func NewAuthRepo(app *forumapp.ForumApp, auth *middlewares.AuthContext) *AuthRepo {
	res := shared.NewJSONRes()
	shared := shared.NewSharedConfig()
	userRepo := NewUserRepo(app.Db.Query)

	// postsService declaration
	postsRepository := postrepo.NewPostRepository(app.Db.Query)
	posts := postrepo.NewPostService(postsRepository)

	userService := NewUserService(userRepo, posts)

	return &AuthRepo{
		app:    app,
		res:    res,
		auth:   auth,
		shared: shared,
		user:   userService,
	}
}
