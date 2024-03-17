package model

import "time"

// Chat ...
type Chat struct {
	Usernames []string `db:"usernames"`
}

// Message ...
type Message struct {
	ChatID    int64
	From      int64
	Text      string
	Timestamp time.Time
}
