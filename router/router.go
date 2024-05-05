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

	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := controllers.NewUserControllers(userUseCase)
	pinpointUseCase := usecase.NewPinpointUseCase(pinpointRepo)
	pinpointController := controllers.NewPinpointControllers(pinpointUseCase)
	missionUseCase := usecase.NewMissionUseCase(missionRepo)
	missionController := controllers.NewMissionControllers(missionUseCase)

	e := echo.New()
	e.GET("/missions", missionController.GetAllMissions)
	e.GET("/missions/:id", missionController.GetMission)
	e.POST("/missions", missionController.CreateMission)
	e.PUT("/missions/:id", missionController.UpdateMission)
	e.DELETE("/missions/:id", missionController.DeleteMission)
	e.GET("/pinpoints", pinpointController.GetAllPinpoints)
	e.GET("/pinpoints/:id", pinpointController.GetPinpoint)
	e.POST("/pinpoints", pinpointController.CreatePinpoint)
	e.PUT("/pinpoints/:id", pinpointController.UpdatePinpoint)
	e.DELETE("/pinpoints/:id", pinpointController.DeletePinpoint)
	e.POST("/users/login", userController.Login)
	e.POST("/users/register", userController.Register)
	e.Logger.Fatal(e.Start(":8080"))

	return *e
}
