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
		response := Response{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Email and password are required",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	token, err := controller.UserUseCase.Login(input.Email, input.Password)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to login",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response := Response{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Login success",
		Token:      token,
	}
	return c.JSON(http.StatusCreated, response)
}

func (controller *UserControllers) Register(c echo.Context) error {
	var input repository.User

	c.Bind(&input)

	if input.Name == "" || input.Email == "" || input.Password == "" {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Name, email, and password are required",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	token, err := controller.UserUseCase.Register(input.Name, input.Email, input.Password, input.Role)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to register",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response := Response{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Register success",
		Token:      token,
	}
	return c.JSON(http.StatusCreated, response)
}
