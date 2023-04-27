package data

import "time"

type Message struct {
	Id          int
	SenderId    int
	RecipientId int
	Content     string
	CreatedAt   time.Time
}

// --- C ---

// --- C End ---

// --- R ---

func (message *Message) GetSender() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", message.SenderId).
		Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
	return
}
func (message *Message) GetRecipient() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", message.RecipientId).
		Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
	return
}
func (message *Message) GetCreatedAtDate() string {
	return message.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// --- R End ---

// --- U ---

// --- U End ---

// --- D ---

// --- D End ---
