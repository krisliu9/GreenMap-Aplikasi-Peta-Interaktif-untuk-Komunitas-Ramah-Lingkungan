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

	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := controllers.NewUserControllers(userUseCase)

	e := echo.New()
	e.POST("/users/login", userController.Login)
	e.POST("/users/register", userController.Register)
	e.Logger.Fatal(e.Start(":8080"))

	return *e
}
