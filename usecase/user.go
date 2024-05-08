package usecase

import (
	"mini-project/auth"
	"mini-project/repository"
)

type UserUseCase struct {
	UserRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepo: repo,
	}
}

func (usecase *UserUseCase) Login(email, password string) (repository.User, string, error) {
	user, err := usecase.UserRepo.Login(email, password)
	if err != nil {
		return user, "", err
	}
	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return user, token, err
	}
	return user, token, nil
}

func (usecase *UserUseCase) Register(name, email, password, role string) (string, error) {
	user, err := usecase.UserRepo.Register(name, email, password, role)
	if err != nil {
		return "", err
	}
	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return token, err
	}
	return token, nil
}
