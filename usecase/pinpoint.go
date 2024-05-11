package usecase

import (
	"mini-project/repository"
)

type PinpointUseCase struct {
	UserMissionUseCase UserMissionUseCase
	PinpointRepo       repository.PinpointRepository
	UserRepo           repository.UserRepository
	MissionUseCase     MissionUseCase
	UserUseCase        UserUseCase
}

func NewPinpointUseCase(userMissionUseCase UserMissionUseCase, pinpointRepo repository.PinpointRepository, userRepo repository.UserRepository, missionUseCase MissionUseCase, userUseCase UserUseCase) *PinpointUseCase {
	return &PinpointUseCase{
		UserMissionUseCase: userMissionUseCase,
		PinpointRepo:       pinpointRepo,
		UserRepo:           userRepo,
		MissionUseCase:     missionUseCase,
		UserUseCase:        userUseCase,
	}
}

func (usecase *PinpointUseCase) GetAllPinpoints() ([]repository.Pinpoint, error) {
	pinpoints, err := usecase.PinpointRepo.GetAll()
	if err != nil {
		return []repository.Pinpoint{}, err
	}
	return pinpoints, nil
}

func (usecase *PinpointUseCase) GetPinpoint(id uint) (repository.Pinpoint, error) {
	pinpoint, err := usecase.PinpointRepo.GetByID(id)
	if err != nil {
		return repository.Pinpoint{}, err
	}
	return pinpoint, nil
}

func (usecase *PinpointUseCase) CreatePinpoint(userId uint, name, description string, latitude, longitude float64) (repository.Pinpoint, error) {
	pinpoint, err := usecase.PinpointRepo.Create(userId, name, description, latitude, longitude)
	if err != nil {
		return repository.Pinpoint{}, err
	}

	currentMissions, err := usecase.UserMissionUseCase.GetMissionByUserId(userId)
	if err != nil {
		return repository.Pinpoint{}, err
	}

	for _, currentMission := range currentMissions {
		userMission, err := usecase.UserMissionUseCase.ProgressMission(currentMission.ID)
		if err != nil {
			return repository.Pinpoint{}, err
		}
		mission, _ := usecase.MissionUseCase.GetMission(currentMission.MissionID)
		if userMission.CurrentProgress == mission.Target {
			usecase.UserRepo.UpdatePoint(userId, mission.Point)
			usecase.UserUseCase.UpdateTier(userId)
		}
	}
	return pinpoint, nil
}

func (usecase *PinpointUseCase) UpdatePinpoint(id uint, name, description string, latitude, longitude float64) (repository.Pinpoint, error) {
	pinpoint, err := usecase.PinpointRepo.Update(id, name, description, latitude, longitude)
	if err != nil {
		return repository.Pinpoint{}, err
	}
	return pinpoint, nil
}

func (usecase *PinpointUseCase) DeletePinpoint(id uint) error {
	err := usecase.PinpointRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
