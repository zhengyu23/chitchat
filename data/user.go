package data

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// --- C ---

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	statement := "insert into users(name, email, password, created_at) " +
		"values($1, $2, $3, $4) returning id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Name, user.Email, user.Password, time.Now()).Scan(&user.Id, &user.CreatedAt)
	return
}

// CreateForum a new thread
func (user *User) CreateForum(title string, content string) (forum Forum, err error) {
	statement := "insert into forums (title, content,author_id, created_at) values($1, $2, $3, $4) returning id, title,content, author_id,created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(title, content, user.Id, time.Now()).Scan(&forum.Id, &forum.Title, &forum.Content, &forum.AuthorId, &forum.CreatedAt)
	return
}

// CreatePost Create a new post to a thread
func (user *User) CreatePost(forum Forum, content string) (replay Replay, err error) {
	statement := "insert into replays (content, author_id, forum_id, created_at) values ($1, $2, $3, $4) returning id, content, author_id, forum_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(content, user.Id, forum.Id, time.Now()).Scan(&replay.Id, &replay.Content, &replay.AuthorId, &replay.ForumId, &replay.CreatedAt)
	return
}

func (user *User) CreateMessage(content string, senderId int, recipientId int) (message Message, err error) {
	statement := "insert into messages (content, sender_id, recipient_id,created_at) values ( $1, $2, $3, $4) returning id, content, sender_id, recipient_id,created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(content, senderId, recipientId, time.Now()).Scan(&message.Id, &message.Content, &message.SenderId, &message.RecipientId, &message.CreatedAt)
	return
}

// --- C End ---

// --- R ---

// UserByEmail Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, name, email, password, created_at from users where email = $1",
		email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (user *User) ReflectDetail() (err error) {
	err = Db.QueryRow("select id, name, email, password, created_at from users where id = $1",
		user.Id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (user *User) Messages() (messages []Message, err error) {
	rows, err := Db.Query("select id, content, sender_id, recipient_id, created_at from messages where sender_id=$1 OR recipient_id=$1 order by created_at DESC", user.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		message := Message{}
		if err = rows.Scan(&message.Id, &message.Content, &message.SenderId, &message.RecipientId, &message.CreatedAt); err != nil {
			return
		}
		messages = append(messages, message)
	}
	rows.Close()
	return
}

func (user *User) GetNumReplies() (count int) {
	row := Db.QueryRow("SELECT count(*) FROM messages where sender_id=$1 OR recipient_id=$1", user.Id)
	err := row.Scan(&count)
	if err != nil {
		return
	}
	return
}

// --- R End ---

// --- U ---

// Update user information in the database
func (user *User) Update() (err error) {
	statement := "update users set name=$2,email=$3 where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Name, user.Email)
	return
}

// --- U End ---

// --- D ---

// Delete user from database
func (user *User) Delete() (err error) {
	statement := "delete from users where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// returns a Result summarizing the effect of the statement.
	_, err = stmt.Exec(user.Id)
	return
}

// --- D End ---
