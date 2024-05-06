package repository

import (
	"gorm.io/gorm"
)

type ReportRepository interface {
	GetAll() ([]Report, error)
	GetByID(id uint) (Report, error)
	Create(reason string) (Report, error)
	Update(id uint, reason string) (Report, error)
	Delete(id uint) error
}

type ReportRepositoryReciever struct {
	DB gorm.DB
}

func NewReportRepository(db gorm.DB) *ReportRepositoryReciever {
	return &ReportRepositoryReciever{
		DB: db,
	}
}

func (r *ReportRepositoryReciever) GetAll() ([]Report, error) {
	var reports []Report
	err := r.DB.Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *ReportRepositoryReciever) GetByID(id uint) (Report, error) {
	var report Report
	err := r.DB.First(&report, id).Error
	if err != nil {
		return Report{}, err
	}
	return report, nil
}

func (r *ReportRepositoryReciever) Create(reason string) (Report, error) {
	report := Report{Reason: reason}
	err := r.DB.Create(&report).Error
	if err != nil {
		return Report{}, err
	}
	return report, nil
}

func (r *ReportRepositoryReciever) Update(id uint, reason string) (Report, error) {
	report, err := r.GetByID(id)
	if err != nil {
		return Report{}, err
	}
	report.Reason = reason
	err = r.DB.Save(&report).Error
	if err != nil {
		return Report{}, err
	}
	return report, nil
}

func (r *ReportRepositoryReciever) Delete(id uint) error {
	report, err := r.GetByID(id)
	if err != nil {
		return err
	}
	err = r.DB.Delete(&report).Error
	if err != nil {
		return err
	}
	return nil
}
