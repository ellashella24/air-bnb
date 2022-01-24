package main

import (
	"air-bnb/configs"
	"air-bnb/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	e := echo.New()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

	fmt.Println(db)
}
