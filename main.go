package main

import (
	"air-bnb/configs"
	controllers "air-bnb/delivery/controllers/homestay"
	"air-bnb/delivery/routes"
	"air-bnb/repository/homestay"
	"air-bnb/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	rp := homestay.NewRepositoryHomestay(db)
	ctrl := controllers.NewControllerHomestay(rp)
	e := echo.New()

	routes.PathHomestay(e, *ctrl)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

}
