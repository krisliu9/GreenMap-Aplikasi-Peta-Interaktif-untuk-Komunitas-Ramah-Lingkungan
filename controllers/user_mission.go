package controllers

import (
	"mini-project/auth"
	"mini-project/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserMissionControllers struct {
	UserMissionUseCase *usecase.UserMissionUseCase
}

func NewUserMissionControllers(missionUseCase *usecase.UserMissionUseCase) *UserMissionControllers {
	return &UserMissionControllers{
		UserMissionUseCase: missionUseCase,
	}
}

type TakeMissionRequest struct {
	MissionId uint `json:"mission_id"`
}

type TakeUserMissionResponse struct {
	ID              uint `json:"id"`
	MissionId       uint `json:"mission_id"`
	UserId          uint `json:"user_id"`
	CurrentProgress int  `json:"current_progress"`
}

func (controller *UserMissionControllers) TakeMission(c echo.Context) error {

	claims, _ := auth.GetTokenClaims(c)
	userId := claims["user_id"]
	var input TakeMissionRequest

	c.Bind(&input)

	if input.MissionId == 0 {
		response := Response{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "All fields must be filled",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	mission, err := controller.UserMissionUseCase.TakeMission(uint(userId.(float64)), input.MissionId)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create mission",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	usermissionResponse := TakeUserMissionResponse{
		ID:              mission.ID,
		MissionId:       mission.MissionID,
		UserId:          mission.UserID,
		CurrentProgress: mission.CurrentProgress,
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Mission created",
		Data:       usermissionResponse,
	}
	return c.JSON(http.StatusCreated, response)
}
