package repository

import (
	"gift-service/internal/models"

	"gorm.io/gorm"
)

type GiftGroupRepo struct {
	db *gorm.DB
}

func NewGiftGroupRepo(db *gorm.DB) *GiftGroupRepo {
	return &GiftGroupRepo{db: db}
}

func (r *GiftGroupRepo) Create(group *models.GiftCardGroup) error {
	return r.db.Create(group).Error
}

func (r *GiftGroupRepo) GetByID(id string) (*models.GiftCardGroup, error) {
	var g models.GiftCardGroup
	if err := r.db.First(&g, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &g, nil
}

func (r *GiftGroupRepo) ListAll() ([]models.GiftCardGroup, error) {
	var list []models.GiftCardGroup
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *GiftGroupRepo) GetStats(groupID string) (total int64, used int64, err error) {
	if err := r.db.Model(&models.GiftCard{}).Where("group_id = ?", groupID).Count(&total).Error; err != nil {
		return 0, 0, err
	}
	if err := r.db.Model(&models.GiftCard{}).Where("group_id = ? AND used = true", groupID).Count(&used).Error; err != nil {
		return 0, 0, err
	}
	return total, used, nil
}
