package data

import "time"

type Replay struct {
	Id        int
	Content   string
	AuthorId  int
	ForumId   int
	CreatedAt time.Time
}

// --- C ---

// --- C End ---

// --- R ---
// User (post *Post) Get the user who wrote the post
func (replay *Replay) GetAuthor() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", replay.AuthorId).
		Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func (replay *Replay) GetCreatedAtDate() string {
	return replay.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// --- R End ---

// --- U ---

// --- U End ---

// --- D ---

// --- D End ---
