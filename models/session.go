package models

import (
	"gorm.io/gorm"
	"time"
)

type ChatSession struct {
	gorm.Model
	ID int64 `gorm:"primaryKey"`

	// Used to check if the chat is inbox/individual or group
	ChatType int

	// Only visible if the session is a group chat
	Name string

	Participants []*User `gorm:"many2many:user_in_chat_sessions"`
}

type UserInChatSessions struct {
	UserID        int64 `gorm:"primaryKey"`
	ChatSessionID int64 `gorm:"primaryKey"`
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
