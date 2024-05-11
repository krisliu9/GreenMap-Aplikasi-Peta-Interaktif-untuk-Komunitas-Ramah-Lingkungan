package router

import (
	"mini-project/controllers"
	"mini-project/database"
	"mini-project/repository"
	"mini-project/usecase"

	"github.com/labstack/echo/v4"
)

func NewRouter() echo.Echo {

	database.InitDatabase()
	userRepo := repository.NewUserRepository(*database.DB)
	pinpointRepo := repository.NewPinpointRepository(*database.DB)
	missionRepo := repository.NewMissionRepository(*database.DB)
	reportRepo := repository.NewReportRepository(*database.DB)
	userMissionRepo := repository.NewUserMissionRepository(*database.DB)
	tierRepo := repository.NewTierRepository(*database.DB)

	tierUseCase := usecase.NewTierUseCase(tierRepo)
	missionUseCase := usecase.NewMissionUseCase(missionRepo)
	userUseCase := usecase.NewUserUseCase(userRepo, tierRepo)
	userMissionUseCase := usecase.NewUserMissionUseCase(userMissionRepo)
	userController := controllers.NewUserControllers(userUseCase)
	pinpointUseCase := usecase.NewPinpointUseCase(*userMissionUseCase, pinpointRepo, userRepo, *missionUseCase, *userUseCase)
	pinpointController := controllers.NewPinpointControllers(pinpointUseCase)
	missionController := controllers.NewMissionControllers(missionUseCase)
	reportUseCase := usecase.NewReportUseCase(reportRepo)
	reportController := controllers.NewReportControllers(reportUseCase)
	userMissionController := controllers.NewUserMissionControllers(userMissionUseCase)

	e := echo.New()
	e.POST("/tiers", controllers.NewTierControllers(tierUseCase).CreateTier)
	e.PUT("/tiers/:id", controllers.NewTierControllers(tierUseCase).UpdateTier)
	e.DELETE("/tiers/:id", controllers.NewTierControllers(tierUseCase).DeleteTier)
	e.GET("/me", userController.Me)
	e.GET("/reports", reportController.GetAllReports)
	e.GET("/reports/:id", reportController.GetReport)
	e.POST("/reports", reportController.CreateReport)
	e.PUT("/reports", reportController.UpdateReport)
	e.DELETE("/reports", reportController.DeleteReport)
	e.GET("/missions", missionController.GetAllMissions)
	e.GET("/missions/:id", missionController.GetMission)
	e.POST("/missions", missionController.CreateMission)
	e.PUT("/missions", missionController.UpdateMission)
	e.DELETE("/missions", missionController.DeleteMission)
	e.GET("/pinpoints", pinpointController.GetAllPinpoints)
	e.GET("/pinpoints/:id", pinpointController.GetPinpoint)
	e.POST("/pinpoints", pinpointController.CreatePinpoint)
	e.PUT("/pinpoints", pinpointController.UpdatePinpoint)
	e.DELETE("/pinpoints", pinpointController.DeletePinpoint)
	e.POST("/users/login", userController.Login)
	e.POST("/users/register", userController.Register)
	e.POST("/usermissions/take", userMissionController.TakeMission)
	e.Logger.Fatal(e.Start(":8080"))

	return *e
}
