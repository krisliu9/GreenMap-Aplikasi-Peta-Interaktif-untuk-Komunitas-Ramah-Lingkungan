package repository

import (
	"time"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetAll() ([]Report, error)
	GetByID(id uint) (Report, error)
	Create(reason string) (Report, error)
	Update(id uint, reason string) (Report, error)
	Delete(id uint) error
}

type ReportRepositoryReceiver struct {
	DB gorm.DB
}

func NewReportRepository(db gorm.DB) *ReportRepositoryReceiver {
	return &ReportRepositoryReceiver{
		DB: db,
	}
}

func (r *ReportRepositoryReceiver) GetAll() ([]Report, error) {
	var reports []Report
	err := r.DB.Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *ReportRepositoryReceiver) GetByID(id uint) (Report, error) {
	var report Report
	err := r.DB.First(&report, id).Error
	if err != nil {
		return Report{}, err
	}
	return report, nil
}

func (r *ReportRepositoryReceiver) Create(reason string) (Report, error) {
	report := Report{Reason: reason}
	err := r.DB.Create(&report).Error
	if err != nil {
		return Report{}, err
	}
	return report, nil
}

func (r *ReportRepositoryReceiver) Update(id uint, reason string) (Report, error) {
	reportUpdate := Report{
		Reason:    reason,
		UpdatedAt: time.Now(),
	}
	if err := r.DB.Model(Report{}).Where("id = ?", id).Updates(&reportUpdate).Error; err != nil {
		return Report{}, err
	}
	return reportUpdate, nil
}

func (r *ReportRepositoryReceiver) Delete(id uint) error {
	err := r.DB.Delete(&Report{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
