package repository

import (
	"time"

	"gift-service/internal/models"

	"gorm.io/gorm"
)

type GiftCardRepo struct {
	db *gorm.DB
}

func NewGiftCardRepo(db *gorm.DB) *GiftCardRepo {
	return &GiftCardRepo{db: db}
}

func (r *GiftCardRepo) CreateMany(cards []*models.GiftCard) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&cards).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *GiftCardRepo) GetByCode(code string) (*models.GiftCard, error) {
	var c models.GiftCard
	if err := r.db.First(&c, "code = ?", code).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *GiftCardRepo) UseCardAtomic(code, phone string) (*models.GiftCard, error) {
	var updated models.GiftCard
	err := r.db.Transaction(func(tx *gorm.DB) error {
		now := time.Now().UTC()
		res := tx.Model(&models.GiftCard{}).
			Where("code = ? AND used = false", code).
			Updates(map[string]interface{}{
				"used":       true,
				"user_phone": phone,
				"used_at":    &now,
			})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return nil
		}
		if err := tx.First(&updated, "code = ?", code).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	// if updated.ID is zero, no update happened
	if updated.ID == (models.GiftCard{}).ID {
		return nil, nil
	}
	return &updated, nil
}

func (r *GiftCardRepo) ListByGroup(groupID string) ([]models.GiftCard, error) {
	var list []models.GiftCard
	if err := r.db.Where("group_id = ?", groupID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *GiftCardRepo) ListUsedByGroup(groupID string) ([]models.GiftCard, error) {
	var list []models.GiftCard
	if err := r.db.Where("group_id = ? AND used = true", groupID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *GiftCardRepo) ListCodesByGroup(groupID string) (used []string, unused []string, err error) {
	var usedCards []models.GiftCard
	var unusedCards []models.GiftCard

	if err := r.db.Where("group_id = ? AND used = ?", groupID, true).Find(&usedCards).Error; err != nil {
		return nil, nil, err
	}
	if err := r.db.Where("group_id = ? AND used = ?", groupID, false).Find(&unusedCards).Error; err != nil {
		return nil, nil, err
	}

	for _, c := range usedCards {
		used = append(used, c.Code)
	}
	for _, c := range unusedCards {
		unused = append(unused, c.Code)
	}

	return used, unused, nil
}

func (r *GiftCardRepo) Update(card *models.GiftCard) error {
	return r.db.Save(card).Error
}
