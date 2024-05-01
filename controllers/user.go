package controllers

import (
	"mini-project/repository"
	"mini-project/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllers struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserControllers(userUseCase *usecase.UserUseCase) *UserControllers {
	return &UserControllers{
		UserUseCase: userUseCase,
	}
}

func (controller *UserControllers) Login(c echo.Context) error {
	var input repository.User

	c.Bind(&input)

	if input.Email == "" || input.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Email and password are required")
	}

	token, err := controller.UserUseCase.Login(input.Email, input.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to login")
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"token": token,
	})
}

func (controller *UserControllers) Register(c echo.Context) error {
	var input repository.User

	c.Bind(&input)

	if input.Name == "" || input.Email == "" || input.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Name, email, and password are required")
	}

	token, err := controller.UserUseCase.Register(input.Name, input.Email, input.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to register")
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"token": token,
	})
}
