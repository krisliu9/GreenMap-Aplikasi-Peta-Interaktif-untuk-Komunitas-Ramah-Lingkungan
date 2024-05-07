package repository

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PinpointRepository interface {
	GetAll() ([]Pinpoint, error)
	GetByID(id uint) (Pinpoint, error)
	Create(name, description string, latitude, longitude float64) (Pinpoint, error)
	Update(id uint, name, description string, latitude, longitude float64) (Pinpoint, error)
	Delete(id uint) error
}

type PinpointRepositoryReceiver struct {
	DB gorm.DB
}

func NewPinpointRepository(db gorm.DB) *PinpointRepositoryReceiver {
	return &PinpointRepositoryReceiver{
		DB: db,
	}
}

func (r *PinpointRepositoryReceiver) GetAll() ([]Pinpoint, error) {
	var pinpoints []Pinpoint
	if err := r.DB.Find(&pinpoints).Error; err != nil {
		return []Pinpoint{}, err
	}
	return pinpoints, nil
}

func (r *PinpointRepositoryReceiver) GetByID(id uint) (Pinpoint, error) {
	var pinpoint Pinpoint
	if err := r.DB.First(&pinpoint, id).Error; err != nil {
		return Pinpoint{}, err
	}
	return pinpoint, nil
}

func (r *PinpointRepositoryReceiver) Create(name, description string, latitude, longitude float64) (Pinpoint, error) {
	pinpoint := Pinpoint{
		Name:        name,
		Description: description,
		Latitude:    latitude,
		Longitude:   longitude,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := r.DB.Create(&pinpoint).Error; err != nil {
		return Pinpoint{}, err
	}
	return pinpoint, nil
}

func (r *PinpointRepositoryReceiver) Update(id uint, name, description string, latitude, longitude float64) (Pinpoint, error) {
	pinpoint := Pinpoint{
		ID:          id,
		Name:        name,
		Description: description,
		Latitude:    latitude,
		Longitude:   longitude,
		UpdatedAt:   time.Now(),
	}
	tx := r.DB.Model(Pinpoint{}).Where("id = ?", id).Updates(pinpoint)
	fmt.Println(tx.Error)
	if tx.Error != nil {
		return Pinpoint{}, tx.Error
	}
	return pinpoint, nil
}

func (r *PinpointRepositoryReceiver) Delete(id uint) error {
	if err := r.DB.Delete(&Pinpoint{}, id).Error; err != nil {
		return err
	}
	return nil
}
