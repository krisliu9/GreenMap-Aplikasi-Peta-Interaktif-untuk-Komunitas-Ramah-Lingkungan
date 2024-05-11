package controllers

import (
	"mini-project/repository"
	"mini-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TierControllers struct {
	TierUseCase *usecase.TierUseCase
}

func NewTierControllers(tierUseCase *usecase.TierUseCase) *TierControllers {
	return &TierControllers{
		TierUseCase: tierUseCase,
	}
}

func (controller *TierControllers) CreateTier(c echo.Context) error {
	var input repository.Tier

	c.Bind(&input)

	if input.Tier_Name == "" || input.Minimal_Point == 0 {
		response := Response{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "All fields must be filled",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	tier, err := controller.TierUseCase.CreateTier(input)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create tier",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Tier created",
		Data:       tier,
	}
	return c.JSON(http.StatusCreated, response)
}

func (controller *TierControllers) UpdateTier(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var input repository.Tier

	c.Bind(&input)

	if input.Tier_Name == "" || input.Minimal_Point == 0 {
		response := Response{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "All fields must be filled",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	tier, err := controller.TierUseCase.UpdateTier(uint(id), input)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update tier",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Tier updated",
		Data:       tier,
	}
	return c.JSON(http.StatusOK, response)
}

func (controller *TierControllers) DeleteTier(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.TierUseCase.DeleteTier(uint(id))
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete tier",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Tier deleted",
	}
	return c.JSON(http.StatusOK, response)
}
