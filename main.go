package main

import (
	"air-bnb/configs"
	tempBookingController "air-bnb/delivery/controllers/booking"
	tempCityController "air-bnb/delivery/controllers/city"
	controllers "air-bnb/delivery/controllers/homestay"
	tempUserController "air-bnb/delivery/controllers/user"
	"air-bnb/delivery/routes"
	tempBookingRepo "air-bnb/repository/booking"
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

	bookingRepo := tempBookingRepo.NewBookingRepository(db)
	bookingController := tempBookingController.NewBookingController(bookingRepo)

	routes.RegisterPath(e, userController, cityController, bookingController)
	rp := homestay.NewRepositoryHomestay(db)
	ctrl := controllers.NewControllerHomestay(rp)

	routes.PathHomestay(e, *ctrl)
	routes.RegisterPath(e, userController, cityController, bc)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

}
