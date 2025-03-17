# Major Database Refactoring

## The first direct database access

The first direct database access is now gone. It was replaced with a new query interface structure.
The first direct concrete implementation was:

commit: `4471caa59b4a4ab9656c44f38a40849e277af378`
Author: adaniel <adiozdaniel@gmail.com>
Date:   Thu Feb 6 14:24:43 2025 +0300

    chores(handlers): add database implementations
    - Add structs for encoding responses and requests.


## The second direct database access

This is a refactored approach of the first direct database access.
At this point, the interfaces, services and seperate concrete implementations are in place but not yet consumed.

commit a14dcb7f47f056a4b2d4cbc149a49ca42e035b51 (HEAD -> bkdIntergration, origin/bkdIntergration)
Author: adioz <adiozdaniel@gmail.com>
Date:   Sat Feb 8 09:27:23 2025 +0300

    refactor(database): change database reference name
    - Changed `forumapp.database.go` database reference to Query.
    - These gives a more readable approach. For instance, to access it: `forumapp.Db.Query` implements a more straight forward approach to reference the queries provided by the database.

## The third approach uses a query interface

The previous approach:

```go
// RegisterHandler handles user registration.
func (h *Repo) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Username == "" {
		h.res.SetError(w, errors.New("username is required"), http.StatusBadRequest)
		return
	}
	if req.Email == "" {
		h.res.SetError(w, errors.New("email is required"), http.StatusBadRequest)
		return
	}
	if req.Password == "" {
		h.res.SetError(w, errors.New("password is required"), http.StatusBadRequest)
		return
	}

	// Check if username or email already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exists int
	err := h.app.Db.Query.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE username = ? OR email = ?", req.Username, req.Email).Scan(&exists)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	if exists > 0 {
		h.res.SetError(w, errors.New("email already exists"), http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	userID, _ := h.app.GenerateUUID()
	// Save to database
	_, err = h.app.Db.Query.ExecContext(ctx, "INSERT INTO users (user_id, username, email, password) VALUES (?, ?, ?, ?)", userID, req.Username, req.Email, string(hashedPassword))
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	// Generate a token (e.g., JWT)
	token := h.app.GenerateToken(userID)

	// Set the session cookie
	http.SetCookie(w, &token)

	// Respond with success and token
	h.res.Err = false
	h.res.Message = "Login successful"
	h.res.Data = map[string]interface{}{
		"token":    token.Value,
		"user_id":  userID,
		"username": req.Username,
	}

	h.res.Err = false
	h.res.Message = "User registered successfully"

	// Respond with JSON
	err = h.res.WriteJSON(w, *h.res, http.StatusCreated)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}
}
```

The query interface basic approach:

```go
// RegisterHandler handles user registration.
func (h *Repo) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Parse user
	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	// Query the database
	err := h.user.Register(&req)
	if err != nil {
		h.res.SetError(w, err, http.StatusConflict)
		return
	}

	// Generate a token (e.g., JWT)
	token := h.app.GenerateToken(req.UserID)

	// Set the session cookie
	http.SetCookie(w, &token)

	// Respond with success and token
	h.res.Data = req
	h.res.Err = false
	h.res.Message = "User registered successfully"

	// Respond with JSON
	err = h.res.WriteJSON(w, *h.res, http.StatusCreated)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}
}
```

Login handler:

```go
// LoginHandler handles user login.
func (h *Repo) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.res.SetError(w, errors.New("invalid request body"), http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Email == "" {
		h.res.SetError(w, errors.New("email is required"), http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		h.res.SetError(w, errors.New("password is required"), http.StatusBadRequest)
		return
	}

	// Fetch user from the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userID, username, hashedPassword string

	err := h.app.Db.Query.QueryRowContext(ctx, "SELECT user_id, username, password FROM users WHERE email = ?", req.Email).
		Scan(&userID, &username, &hashedPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.res.SetError(w, errors.New("user does not exist"), http.StatusUnauthorized)
			return
		}

		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		h.res.SetError(w, errors.New("invalid email or password"), http.StatusUnauthorized)
		return
	}

	// Check if user already has a valid session
	if cookie, err := r.Cookie("session_token"); err == nil {
		if storedCookie, exists := h.app.Sessions.Load(userID); exists {
			if token, ok := storedCookie.(*http.Cookie); ok && token.Value == cookie.Value {
				h.res.Err = false
				h.res.Message = "Login successful (existing session)"
				h.res.Data = map[string]interface{}{
					"token":    token.Value,
					"user_id":  userID,
					"username": username,
				}
				h.res.WriteJSON(w, *h.res, http.StatusOK)
				return
			}
		}
	}

	// Generate a token (e.g., JWT)
	token := h.app.GenerateToken(userID)

	// Set the session cookie
	http.SetCookie(w, &token)

	// Respond with success and token
	h.res.Err = false
	h.res.Message = "Login successful"
	h.res.Data = map[string]interface{}{
		"token":    token.Value,
		"user_id":  userID,
		"username": username,
	}

	err = h.res.WriteJSON(w, *h.res, http.StatusOK)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}
}
```

The basic interface query approach:

```go
// LoginHandler handles user login.
func (h *Repo) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.res.SetError(w, errors.New("invalid request body"), http.StatusBadRequest)
		return
	}

	user, err := h.user.Login(req.Email, req.Password)
	if err != nil {
		h.res.SetError(w, err, http.StatusUnauthorized)
		return
	}

	// Check if user already has a valid session
	if cookie, err := r.Cookie("session_token"); err == nil {
		if storedCookie, exists := h.app.Sessions.Load(user.UserID); exists {
			if token, ok := storedCookie.(*http.Cookie); ok && token.Value == cookie.Value {
				h.res.Err = false
				h.res.Message = "Login successful (existing session)"
				h.res.Data = &user
				h.res.WriteJSON(w, *h.res, http.StatusOK)
				return
			}
		}
	}

	// Generate a token (e.g., JWT)
	token := h.app.GenerateToken(user.UserID)

	// Set the session cookie
	http.SetCookie(w, &token)

	// Respond with success and token
	h.res.Err = false
	h.res.Message = "Login successful"
	h.res.Data = &user

	err = h.res.WriteJSON(w, *h.res, http.StatusOK)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}
}

```