package data

import "time"

type Forum struct {
	Id        int
	Title     string
	Content   string
	AuthorId  int
	CreatedAt time.Time
}

// --- C ---

// --- C End ---

// --- R ---

func Forums() (forums []Forum, err error) {
	rows, err := Db.Query("select id, title, content, author_id, created_at from forums order by created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		forum := Forum{}
		if err = rows.Scan(&forum.Id, &forum.Title, &forum.Content, &forum.AuthorId, &forum.CreatedAt); err != nil {
			return
		}
		forums = append(forums, forum)
	}
	rows.Close()
	return
}
func ForumById(id string) (forum Forum, err error) {
	forum = Forum{}
	err = Db.QueryRow("select id,title,content, author_id,created_at from forums where id = $1", id).Scan(&forum.Id, &forum.Title, &forum.Content, &forum.AuthorId, &forum.CreatedAt)
	return
}

// GetAuthor Get the user who started this thread
func (forum *Forum) GetAuthor() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", forum.AuthorId).
		Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func (forum *Forum) GetCreatedAtDate() string {
	return forum.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// GetNumReplies get the number of posts in a thread
func (forum *Forum) GetNumReplies() (count int) {
	row := Db.QueryRow("SELECT count(*) FROM replays where forum_id = $1", forum.Id)
	err := row.Scan(&count)
	if err != nil {
		return
	}
	return
}

// Posts get posts to a thread
func (forum *Forum) GetReplays() (replays []Replay, err error) {
	rows, err := Db.Query("select id, content, author_id, forum_id, created_at from replays where forum_id = $1", forum.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		replay := Replay{}
		if err = rows.Scan(&replay.Id, &replay.Content, &replay.AuthorId, &replay.ForumId, &replay.CreatedAt); err != nil {
			return
		}
		replays = append(replays, replay)
	}
	rows.Close()
	return
}

// --- R End ---

// --- U ---

// --- U End ---

// --- D ---

// --- D End ---
