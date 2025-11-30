package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	WalletID  uuid.UUID `gorm:"type:uuid;index;not null" json:"wallet_id"`
	Amount    int64     `gorm:"not null" json:"amount"`
	Reference string    `gorm:"size:255" json:"reference"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}
