package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GiftCardGroup struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null;index" json:"name"`
	AmountDue int64     `gorm:"not null" json:"amount_due"`
	CreatedAt time.Time `json:"created_at"`
}

func (g *GiftCardGroup) BeforeCreate(tx *gorm.DB) (err error) {
	if g.ID == uuid.Nil {
		g.ID = uuid.New()
	}
	return nil
}
