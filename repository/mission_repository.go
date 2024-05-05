package repository

import (
	"time"

	"gorm.io/gorm"
)

type MissionRepository interface {
	GetAll() ([]Mission, error)
	GetByID(id uint) (Mission, error)
	Create(target int, description string, point int, startAt, endAt time.Time) (Mission, error)
	Update(id uint, target int, description string, point int, startAt, endAt time.Time) (Mission, error)
	Delete(id uint) error
}

type MissionRepositoryReciever struct {
	DB gorm.DB
}

func NewMissionRepository(db gorm.DB) *MissionRepositoryReciever {
	return &MissionRepositoryReciever{
		DB: db,
	}
}

func (r *MissionRepositoryReciever) GetAll() ([]Mission, error) {
	var missions []Mission
	if err := r.DB.Find(&missions).Error; err != nil {
		return []Mission{}, err
	}
	return missions, nil
}

func (r *MissionRepositoryReciever) GetByID(id uint) (Mission, error) {
	var mission Mission
	if err := r.DB.First(&mission, id).Error; err != nil {
		return Mission{}, err
	}
	return mission, nil
}

func (r *MissionRepositoryReciever) Create(target int, description string, point int, startAt, endAt time.Time) (Mission, error) {
	mission := Mission{
		Target:      target,
		Description: description,
		Point:       point,
		StartAt:     startAt,
		EndAt:       endAt,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := r.DB.Create(&mission).Error; err != nil {
		return Mission{}, err
	}
	return mission, nil
}

func (r *MissionRepositoryReciever) Update(id uint, target int, description string, point int, startAt, endAt time.Time) (Mission, error) {
	mission := Mission{
		Target:      target,
		Description: description,
		Point:       point,
		StartAt:     startAt,
		EndAt:       endAt,
		UpdatedAt:   time.Now(),
	}
	if err := r.DB.Model(&mission).Where("id = ?", id).Updates(&mission).Error; err != nil {
		return Mission{}, err
	}
	return mission, nil
}

func (r *MissionRepositoryReciever) Delete(id uint) error {
	if err := r.DB.Where("id = ?", id).Delete(&Mission{}).Error; err != nil {
		return err
	}
	return nil
}
