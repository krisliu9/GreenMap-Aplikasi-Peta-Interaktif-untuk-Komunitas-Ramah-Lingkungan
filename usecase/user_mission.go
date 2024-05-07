package usecase

import (
	"mini-project/repository"
)

type UserMissionUseCase struct {
	UserMissionRepo repository.UserMissionRepository
}

func NewUserMissionUseCase(repo repository.UserMissionRepository) *UserMissionUseCase {
	return &UserMissionUseCase{
		UserMissionRepo: repo,
	}
}

func (usecase *UserMissionUseCase) TakeMission(userId, missionId uint) (repository.UserMission, error) {
	userMission, err := usecase.UserMissionRepo.Create(userId, missionId)
	if err != nil {
		return userMission, err
	}
	return userMission, nil
}

func (usecase *UserMissionUseCase) GetMissionByUserId(userId uint) ([]repository.UserMission, error) {
	userMission, err := usecase.UserMissionRepo.GetAllByUserID(userId)
	if err != nil {
		return userMission, err
	}
	return userMission, nil
}

func (usecase *UserMissionUseCase) ProgressMission(userMissionId uint) (repository.UserMission, error) {
	userMission, err := usecase.UserMissionRepo.Update(userMissionId)
	if err != nil {
		return userMission, err
	}
	return userMission, nil
}
