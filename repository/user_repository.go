package repository

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(email, password string) (User, error)
	Register(name, email, password, role string) (User, error)
	UpdatePoint(userId uint, point int) (User, error)
	UpdateTier(userId uint, tierId uint) (User, error)
	GetByID(userId uint) (User, error)
}

type UserRepositoryReceiver struct {
	DB gorm.DB
}

func NewUserRepository(db gorm.DB) *UserRepositoryReceiver {
	return &UserRepositoryReceiver{
		DB: db,
	}
}

func (repo *UserRepositoryReceiver) Login(email, password string) (User, error) {
	var user User
	repo.DB = *repo.DB.Debug()
	result := repo.DB.First(&user, "email = ? AND password = ?", email, password)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepositoryReceiver) Register(name, email, password, role string) (User, error) {
	user := User{Name: name, Email: email, Password: password, Role: role, Tier_ID: 1}
	result := repo.DB.Omit("Tier_Name").Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (repo *UserRepositoryReceiver) GetByID(userId uint) (User, error) {
	var user User
	if err := repo.DB.First(&user, userId).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (repo *UserRepositoryReceiver) UpdatePoint(userId uint, point int) (User, error) {
	user, err := repo.GetByID(userId)
	if err != nil {
		return User{}, err
	}
	userUpdate := User{
		Current_Point: user.Current_Point + point,
		UpdatedAt:     time.Now(),
	}
	if err := repo.DB.Model(User{}).Where("id = ?", userId).Updates(&userUpdate).Error; err != nil {
		return User{}, err
	}
	return userUpdate, nil
}

func (repo *UserRepositoryReceiver) UpdateTier(userId, tierId uint) (User, error) {
	userUpdate := User{
		Tier_ID:   tierId,
		UpdatedAt: time.Now(),
	}
	if err := repo.DB.Model(User{}).Where("id = ?", userId).Updates(&userUpdate).Error; err != nil {
		return User{}, err
	}
	return userUpdate, nil
}
