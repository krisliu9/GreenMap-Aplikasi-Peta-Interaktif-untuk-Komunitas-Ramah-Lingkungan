package usecase

import (
	"mini-project/repository"
)

type ReportUseCase struct {
	ReportRepo repository.ReportRepository
}

func NewReportUseCase(repo repository.ReportRepository) *ReportUseCase {
	return &ReportUseCase{
		ReportRepo: repo,
	}
}

func (usecase *ReportUseCase) GetAllReports() ([]repository.Report, error) {
	reports, err := usecase.ReportRepo.GetAll()
	if err != nil {
		return []repository.Report{}, err
	}
	return reports, nil
}

func (usecase *ReportUseCase) GetReport(id uint) (repository.Report, error) {
	report, err := usecase.ReportRepo.GetByID(id)
	if err != nil {
		return repository.Report{}, err
	}
	return report, nil
}

func (usecase *ReportUseCase) CreateReport(reason string) (repository.Report, error) {
	report, err := usecase.ReportRepo.Create(reason)
	if err != nil {
		return repository.Report{}, err
	}
	return report, nil
}

func (usecase *ReportUseCase) UpdateReport(id uint, reason string) (repository.Report, error) {
	report, err := usecase.ReportRepo.Update(id, reason)
	if err != nil {
		return repository.Report{}, err
	}
	return report, nil
}

func (usecase *ReportUseCase) DeleteReport(id uint) error {
	err := usecase.ReportRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
