package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GiftCard struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	GroupID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"group_id"`
	Code      string     `gorm:"size:64;uniqueIndex;not null" json:"code"`
	Amount    int64      `gorm:"not null" json:"amount"`
	Used      bool       `gorm:"not null;default:false" json:"used"`
	UserPhone string     `gorm:"size:32" json:"user_phone,omitempty"`
	UsedAt    *time.Time `json:"used_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

func (g *GiftCard) BeforeCreate(tx *gorm.DB) (err error) {
	if g.ID == uuid.Nil {
		g.ID = uuid.New()
	}
	return nil
}
