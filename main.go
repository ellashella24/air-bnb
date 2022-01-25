package main

import (
	"air-bnb/configs"
	tempCityController "air-bnb/delivery/controllers/city"
	controllers "air-bnb/delivery/controllers/homestay"
	tempUserController "air-bnb/delivery/controllers/user"
	"air-bnb/delivery/routes"
	tempCityRepo "air-bnb/repository/city"
	"air-bnb/repository/homestay"
	tempUserRepo "air-bnb/repository/user"
	"air-bnb/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	e := echo.New()

	userRepo := tempUserRepo.NewUserRepository(db)
	userController := tempUserController.NewUserController(userRepo)

	cityRepo := tempCityRepo.NewCityRepository(db)
	cityController := tempCityController.NewCityController(cityRepo)

	rp := homestay.NewRepositoryHomestay(db)
	ctrl := controllers.NewControllerHomestay(rp)

	routes.PathHomestay(e, *ctrl)
	routes.RegisterPath(e, userController, cityController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

}
