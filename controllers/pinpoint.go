package controllers

import (
	"mini-project/auth"
	"mini-project/repository"
	"mini-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PinpointControllers struct {
	PinpointUseCase *usecase.PinpointUseCase
}

func NewPinpointControllers(pinpointUseCase *usecase.PinpointUseCase) *PinpointControllers {
	return &PinpointControllers{
		PinpointUseCase: pinpointUseCase,
	}
}

func (controller *PinpointControllers) GetAllPinpoints(c echo.Context) error {
	pinpoints, err := controller.PinpointUseCase.GetAllPinpoints()
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get all pinpoints",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "All pinpoints",
		Data:       pinpoints,
	}
	return c.JSON(http.StatusOK, response)
}

func (controller *PinpointControllers) GetPinpoint(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	pinpoint, err := controller.PinpointUseCase.GetPinpoint(uint(id))
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusNotFound,
			Message:    "Pinpoint not found",
		}
		return echo.NewHTTPError(http.StatusNotFound, response)
	}
	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Pinpoint found",
		Data:       pinpoint,
	}
	return c.JSON(http.StatusOK, response)
}

type PinpointInsertResponse struct {
	ID             uint    `json:"id"`
	UserID         uint    `json:"user_id"`
	PinpointTypeID uint    `json:"pinpoint_type_id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}

func (controller *PinpointControllers) CreatePinpoint(c echo.Context) error {
	claims, _ := auth.GetTokenClaims(c)
	userId := claims["user_id"]
	var input repository.Pinpoint

	c.Bind(&input)

	if input.Name == "" || input.Description == "" || input.Latitude == 0 || input.Longitude == 0 {
		response := Response{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Name, description, latitude, and longitude are required",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	pinpoint, err := controller.PinpointUseCase.CreatePinpoint(uint(userId.(float64)), input.Name, input.Description, input.Latitude, input.Longitude)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create pinpoint",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	pinpointResponse := PinpointInsertResponse{
		ID:     pinpoint.ID,
		UserID: pinpoint.UserID,
		// PinpointTypeID: pinpoint.PinpointTypeID,
		Name:        pinpoint.Name,
		Description: pinpoint.Description,
		Latitude:    pinpoint.Latitude,
		Longitude:   pinpoint.Longitude,
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Pinpoint created",
		Data:       pinpointResponse,
	}
	return c.JSON(http.StatusCreated, response)
}

func (controller *PinpointControllers) UpdatePinpoint(c echo.Context) error {
	var input repository.Pinpoint

	c.Bind(&input)

	if input.Name == "" || input.Description == "" || input.Latitude == 0 || input.Longitude == 0 {
		response := Response{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Name, description, latitude, and longitude are required",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	pinpoint, err := controller.PinpointUseCase.UpdatePinpoint(uint(input.ID), input.Name, input.Description, input.Latitude, input.Longitude)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update pinpoint",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	pinpointResponse := PinpointInsertResponse{
		ID:     pinpoint.ID,
		UserID: pinpoint.UserID,
		// PinpointTypeID: pinpoint.PinpointTypeID,
		Name:        pinpoint.Name,
		Description: pinpoint.Description,
		Latitude:    pinpoint.Latitude,
		Longitude:   pinpoint.Longitude,
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Pinpoint updated",
		Data:       pinpointResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func (controller *PinpointControllers) DeletePinpoint(c echo.Context) error {
	var input repository.Pinpoint
	c.Bind(&input)

	err := controller.PinpointUseCase.DeletePinpoint(uint(input.ID))
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete pinpoint",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Pinpoint deleted",
	}
	return c.JSON(http.StatusOK, response)
}
