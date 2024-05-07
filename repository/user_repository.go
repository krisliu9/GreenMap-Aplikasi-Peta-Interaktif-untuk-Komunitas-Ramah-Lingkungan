package repository

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(email, password string) (User, error)
	Register(name, email, password, role string) (User, error)
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
	result := repo.DB.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepositoryReceiver) Register(name, email, password, role string) (User, error) {
	user := User{Name: name, Email: email, Password: password, Role: role}
	result := repo.DB.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
