package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID int64 `gorm:"primaryKey"`

	// User handle or unique username
	Handle string `gorm:"unique"`

	// User's real name
	Name string

	Chats []ChatSession `gorm:"many2many:user_in_chat_sessions"`
}
