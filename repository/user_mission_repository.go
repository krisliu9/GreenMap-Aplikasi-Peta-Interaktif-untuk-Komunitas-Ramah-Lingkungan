package repository

import (
	"time"

	"gorm.io/gorm"
)

type UserMissionRepository interface {
	GetAll() ([]UserMission, error)
	GetAllByUserID(userId uint) ([]UserMission, error)
	GetByID(id uint) (UserMission, error)
	Create(userID, missionID uint) (UserMission, error)
	Update(id uint) (UserMission, error)
}

type UserMissionRepositoryReceiver struct {
	DB gorm.DB
}

func NewUserMissionRepository(db gorm.DB) *UserMissionRepositoryReceiver {
	return &UserMissionRepositoryReceiver{
		DB: db,
	}
}

func (r *UserMissionRepositoryReceiver) GetAll() ([]UserMission, error) {
	var userMissions []UserMission
	if err := r.DB.Find(&userMissions).Error; err != nil {
		return []UserMission{}, err
	}
	return userMissions, nil
}

func (r *UserMissionRepositoryReceiver) GetAllByUserID(userId uint) ([]UserMission, error) {
	var userMissions []UserMission
	if err := r.DB.Find(&userMissions).Where("user_id = ?", userId).Error; err != nil {
		return []UserMission{}, err
	}
	return userMissions, nil
}

func (r *UserMissionRepositoryReceiver) GetByID(id uint) (UserMission, error) {
	var userMission UserMission
	if err := r.DB.First(&userMission, id).Error; err != nil {
		return UserMission{}, err
	}
	return userMission, nil
}

func (r *UserMissionRepositoryReceiver) Create(userId, missionId uint) (UserMission, error) {
	userMission := UserMission{
		UserID:          userId,
		MissionID:       missionId,
		CurrentProgress: 0,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	if err := r.DB.Create(&userMission).Error; err != nil {
		return UserMission{}, err
	}
	return userMission, nil
}

func (r *UserMissionRepositoryReceiver) Update(id uint) (UserMission, error) {
	userMission, err := r.GetByID(id)
	if err != nil {
		return userMission, err
	}
	userMissionUpdate := UserMission{
		CurrentProgress: userMission.CurrentProgress + 1,
		UpdatedAt:       time.Now(),
	}
	if err := r.DB.Model(UserMission{}).Where("id = ?", id).Updates(&userMissionUpdate).Error; err != nil {
		return UserMission{}, err
	}
	return userMission, nil
}
