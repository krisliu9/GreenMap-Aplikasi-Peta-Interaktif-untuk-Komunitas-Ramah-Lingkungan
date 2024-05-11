package repository

import (
	"gorm.io/gorm"
)

type TierRepository interface {
	GetAllTier() ([]Tier, error)
	GetTierByID(tierId uint) (Tier, error)
	CreateTier(tier Tier) (Tier, error)
	UpdateTier(tierId uint, tier Tier) (Tier, error)
	DeleteTier(tierId uint) error
}

type TierRepositoryReceiver struct {
	DB gorm.DB
}

func NewTierRepository(db gorm.DB) *TierRepositoryReceiver {
	return &TierRepositoryReceiver{
		DB: db,
	}
}

func (repo *TierRepositoryReceiver) GetAllTier() ([]Tier, error) {
	var tier []Tier
	if err := repo.DB.Find(&tier).Error; err != nil {
		return []Tier{}, err
	}
	return tier, nil
}

func (repo *TierRepositoryReceiver) GetTierByID(tierId uint) (Tier, error) {
	var tier Tier
	if err := repo.DB.First(&tier, tierId).Error; err != nil {
		return Tier{}, err
	}
	return tier, nil
}

func (repo *TierRepositoryReceiver) CreateTier(tier Tier) (Tier, error) {
	result := repo.DB.Create(&tier)
	if result.Error != nil {
		return tier, result.Error
	}
	return tier, nil
}

func (repo *TierRepositoryReceiver) UpdateTier(tierId uint, tier Tier) (Tier, error) {
	result := repo.DB.Model(&Tier{}).Where("id = ?", tierId).Updates(tier)
	if result.Error != nil {
		return tier, result.Error
	}
	return tier, nil
}

func (repo *TierRepositoryReceiver) DeleteTier(tierId uint) error {
	result := repo.DB.Delete(&Tier{}, tierId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
