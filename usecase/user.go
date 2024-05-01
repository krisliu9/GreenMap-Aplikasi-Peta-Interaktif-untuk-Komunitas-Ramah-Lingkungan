package usecase

import (
	"mini-project/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserUseCase struct {
	UserRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepo: repo,
	}
}

func (usecase *UserUseCase) Login(email, password string) (string, error) {
	user, err := usecase.UserRepo.Login(email, password)
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	t, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (usecase *UserUseCase) Register(name, email, password string) (string, error) {
	user, err := usecase.UserRepo.Register(name, email, password)
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	t, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}
	return t, nil
}
