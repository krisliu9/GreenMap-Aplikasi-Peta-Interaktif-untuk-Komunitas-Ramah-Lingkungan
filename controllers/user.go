package controllers

import (
	"mini-project/auth"
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

type UserInsertResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	Current_Point int    `json:"current_point"`
	Tier_ID       uint   `json:"tier_id,omitempty"`
	Tier_Name     string `json:"tier_name"`
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

	input, token, err := controller.UserUseCase.Login(input.Email, input.Password)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to login",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	userResponse := UserInsertResponse{
		ID:            input.ID,
		Name:          input.Name,
		Email:         input.Email,
		Role:          input.Role,
		Current_Point: input.Current_Point,
		Tier_ID:       input.Tier_ID,
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Login success",
		Token:      token,
		Data:       userResponse,
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
	userResponse := UserInsertResponse{
		ID:            input.ID,
		Name:          input.Name,
		Email:         input.Email,
		Role:          input.Role,
		Current_Point: input.Current_Point,
		Tier_ID:       input.Tier_ID,
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Register success",
		Token:      token,
		Data:       userResponse,
	}
	return c.JSON(http.StatusCreated, response)
}

func (controller *UserControllers) Me(c echo.Context) error {
	claims, _ := auth.GetTokenClaims(c)
	userId := claims["user_id"]

	user, err := controller.UserUseCase.GetByID(uint(userId.(float64)))
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get user",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	userResponse := UserInsertResponse{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		Role:          user.Role,
		Current_Point: user.Current_Point,
		Tier_Name:     user.Tier_Name,
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Success get user",
		Data:       userResponse,
	}
	return c.JSON(http.StatusOK, response)
}
