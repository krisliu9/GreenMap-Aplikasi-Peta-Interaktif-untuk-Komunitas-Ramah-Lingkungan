package usecase

import (
	"mini-project/repository"
	"time"
)

type MissionUseCase struct {
	MissionRepo repository.MissionRepository
}

func NewMissionUseCase(repo repository.MissionRepository) *MissionUseCase {
	return &MissionUseCase{
		MissionRepo: repo,
	}
}

func (usecase *MissionUseCase) GetAllMissions() ([]repository.Mission, error) {
	missions, err := usecase.MissionRepo.GetAll()
	if err != nil {
		return []repository.Mission{}, err
	}
	return missions, nil
}

func (usecase *MissionUseCase) GetMission(id uint) (repository.Mission, error) {
	mission, err := usecase.MissionRepo.GetByID(id)
	if err != nil {
		return repository.Mission{}, err
	}
	return mission, nil
}

func (usecase *MissionUseCase) CreateMission(target int, description string, point int, startAt, endAt string) (repository.Mission, error) {
	startAtTime, err := time.Parse("02-01-2006 15:04:05 -07", startAt)
	if err != nil {
		return repository.Mission{}, err
	}
	endAtTime, err := time.Parse("02-01-2006 15:04:05 -07", endAt)
	if err != nil {
		return repository.Mission{}, err
	}
	mission, err := usecase.MissionRepo.Create(target, description, point, startAtTime, endAtTime)
	if err != nil {
		return repository.Mission{}, err
	}
	return mission, nil
}

func (usecase *MissionUseCase) UpdateMission(id uint, target int, description string, point int, startAt, endAt string) (repository.Mission, error) {
	startAtTime, err := time.Parse("02-01-2006 15:04:05 -07", startAt)
	if err != nil {
		return repository.Mission{}, err
	}
	endAtTime, err := time.Parse("02-01-2006 15:04:05 -07", endAt)
	if err != nil {
		return repository.Mission{}, err
	}
	mission, err := usecase.MissionRepo.Update(id, target, description, point, startAtTime, endAtTime)
	if err != nil {
		return repository.Mission{}, err
	}
	return mission, nil
}

func (usecase *MissionUseCase) DeleteMission(id uint) error {
	err := usecase.MissionRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
