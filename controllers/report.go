package controllers

import (
	"mini-project/repository"
	"mini-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReportControllers struct {
	ReportUseCase *usecase.ReportUseCase
}

func NewReportControllers(reportUseCase *usecase.ReportUseCase) *ReportControllers {
	return &ReportControllers{
		ReportUseCase: reportUseCase,
	}
}

func (controller *ReportControllers) GetAllReports(c echo.Context) error {
	reports, err := controller.ReportUseCase.GetAllReports()
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get all reports",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}
	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "All reports",
		Data:       reports,
	}
	return c.JSON(http.StatusOK, response)
}

func (controller *ReportControllers) GetReport(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	report, err := controller.ReportUseCase.GetReport(uint(id))
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusNotFound,
			Message:    "Report not found",
		}
		return echo.NewHTTPError(http.StatusNotFound, response)
	}
	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Report found",
		Data:       report,
	}
	return c.JSON(http.StatusOK, response)
}

type ReportInsertResponse struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	PinpointID uint   `json:"pinpoint_id"`
	Reason     string `json:"reason"`
}

func (controller *ReportControllers) CreateReport(c echo.Context) error {
	var input repository.Report
	c.Bind(&input)

	if input.Reason == "" {
		response := Response{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Reason is required",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	report, err := controller.ReportUseCase.CreateReport(input.Reason)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create report",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	reportResponse := ReportInsertResponse{
		ID:         report.ID,
		UserID:     report.UserID,
		PinpointID: report.PinpointID,
		Reason:     report.Reason,
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Report created",
		Data:       reportResponse,
	}
	return c.JSON(http.StatusCreated, response)
}

func (controller *ReportControllers) UpdateReport(c echo.Context) error {
	var input repository.Report
	c.Bind(&input)

	if input.Reason == "" {
		response := Response{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Reason is required",
		}
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	report, err := controller.ReportUseCase.UpdateReport(uint(input.ID), input.Reason)
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update report",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	reportResponse := ReportInsertResponse{
		ID:         report.ID,
		UserID:     report.UserID,
		PinpointID: report.PinpointID,
		Reason:     report.Reason,
	}

	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Report updated",
		Data:       reportResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func (controller *ReportControllers) DeleteReport(c echo.Context) error {
	var input repository.Report
	c.Bind(&input)

	err := controller.ReportUseCase.DeleteReport(uint(input.ID))
	if err != nil {
		response := Response{
			Status:     false,
			StatusCode: http.StatusNotFound,
			Message:    "Report not found",
		}
		return echo.NewHTTPError(http.StatusNotFound, response)
	}
	response := Response{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Report deleted",
	}
	return c.JSON(http.StatusOK, response)
}
