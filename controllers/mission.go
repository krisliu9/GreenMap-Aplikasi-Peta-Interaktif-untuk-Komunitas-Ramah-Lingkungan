package controllers

import (
	"mini-project/usecase"
	"net/http"
	"strconv"

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
	Target      int    `json:"target"`
	Description string `json:"description"`
	Point       int    `json:"point"`
	StartAt     string `json:"start_at"`
	EndAt       string `json:"end_at"`
}

func (controller *MissionControllers) CreateMission(c echo.Context) error {
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

	response := Response{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Mission created",
		Data:       mission,
	}
	return c.JSON(http.StatusCreated, response)
}

func (controller *MissionControllers) UpdateMission(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
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

	mission, err := controller.MissionUseCase.UpdateMission(uint(id), input.Target, input.Description, input.Point, input.StartAt, input.EndAt)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update mission",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Mission updated",
		Data:       mission,
	}
	return c.JSON(http.StatusOK, response)
}

func (controller *MissionControllers) DeleteMission(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.MissionUseCase.DeleteMission(uint(id))
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
