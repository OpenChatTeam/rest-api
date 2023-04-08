package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID        int64 `gorm:"primaryKey"`
	ChannelID int64
	Content   string
}
