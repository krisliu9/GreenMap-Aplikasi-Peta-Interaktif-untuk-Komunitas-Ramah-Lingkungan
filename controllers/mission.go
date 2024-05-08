package controllers

import (
	"fmt"
	"mini-project/auth"
	"mini-project/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type MissionControllers struct {
	MissionUseCase *usecase.MissionUseCase
}

func NewMissionControllers(missionUseCase *usecase.MissionUseCase) *MissionControllers {
	return &MissionControllers{
		MissionUseCase: missionUseCase,
	}
}

func (controller *MissionControllers) GetAllMissions(c echo.Context) error {
	missions, err := controller.MissionUseCase.GetAllMissions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, missions)
}

func (controller *MissionControllers) GetMission(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	mission, err := controller.MissionUseCase.GetMission(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, mission)
}

type MissionInsertRequest struct {
	ID          int    `json:"id"`
	Target      int    `json:"target"`
	Description string `json:"description"`
	Point       int    `json:"point"`
	StartAt     string `json:"start_at"`
	EndAt       string `json:"end_at"`
}

type MissionInsertResponse struct {
	ID          uint      `json:"id"`
	Target      int       `json:"target"`
	Description string    `json:"description"`
	Point       int       `json:"point"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
}

func (controller *MissionControllers) CreateMission(c echo.Context) error {
	claims, err := auth.GetTokenClaims(c)
	role := claims["role"]
	if err != nil || role != auth.RoleAdmin {
		fmt.Println(err)
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    http.StatusText(http.StatusUnauthorized),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	var input MissionInsertRequest

	c.Bind(&input)

	if input.Target == 0 || input.Description == "" || input.Point == 0 || input.StartAt == "" || input.EndAt == "" {
		response := Response{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "All fields must be filled",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	mission, err := controller.MissionUseCase.CreateMission(input.Target, input.Description, input.Point, input.StartAt, input.EndAt)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create mission",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	missionResponse := MissionInsertResponse{
		ID:          mission.ID,
		Target:      mission.Target,
		Description: mission.Description,
		Point:       mission.Point,
		StartAt:     mission.StartAt,
		EndAt:       mission.EndAt,
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Mission created",
		Data:       missionResponse,
	}
	return c.JSON(http.StatusCreated, response)
}

func (controller *MissionControllers) UpdateMission(c echo.Context) error {
	claims, err := auth.GetTokenClaims(c)
	role := claims["role"]
	if err != nil || role != auth.RoleAdmin {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    http.StatusText(http.StatusUnauthorized),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	var input MissionInsertRequest

	c.Bind(&input)

	if input.Target == 0 || input.Description == "" || input.Point == 0 || input.StartAt == "" || input.EndAt == "" {
		response := Response{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "All fields must be filled",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	mission, err := controller.MissionUseCase.UpdateMission(uint(input.ID), input.Target, input.Description, input.Point, input.StartAt, input.EndAt)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update mission",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	missionResponse := MissionInsertResponse{
		ID:          mission.ID,
		Target:      mission.Target,
		Description: mission.Description,
		Point:       mission.Point,
		StartAt:     mission.StartAt,
		EndAt:       mission.EndAt,
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Mission updated",
		Data:       missionResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func (controller *MissionControllers) DeleteMission(c echo.Context) error {
	claims, err := auth.GetTokenClaims(c)
	role := claims["role"]
	if err != nil || role != auth.RoleAdmin {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    http.StatusText(http.StatusUnauthorized),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	var input MissionInsertRequest
	c.Bind(&input)

	err = controller.MissionUseCase.DeleteMission(uint(input.ID))
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete mission",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Mission deleted",
	}
	return c.JSON(http.StatusOK, response)
}
