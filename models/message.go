package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	MessageId pgtype.Numeric
	ChannelId pgtype.Numeric
	Content   string
}
