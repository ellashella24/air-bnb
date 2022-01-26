package main

import (
	"air-bnb/configs"
	tempBookingController "air-bnb/delivery/controllers/booking"
	tempCityController "air-bnb/delivery/controllers/city"
	tempUserController "air-bnb/delivery/controllers/user"
	"air-bnb/delivery/routes"
	tempBookingRepo "air-bnb/repository/booking"
	tempCityRepo "air-bnb/repository/city"
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

	bookingRepo := tempBookingRepo.NewBookingRepository(db)
	tempBookingController := tempBookingController.NewBookingController(bookingRepo)

	routes.RegisterPath(e, userController, cityController, tempBookingController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

	fmt.Println(db)
}
