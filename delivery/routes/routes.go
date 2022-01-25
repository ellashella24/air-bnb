package routes

import (
	"air-bnb/constants"
	controllers "air-bnb/delivery/controllers/homestay"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func PathHomestay(e *echo.Echo, c controllers.StructCtrlHomestay) {

	e.GET("/homestays", c.GetAllHomestay())
	e.GET("/homestays/host", c.GetHostHomestay(), middleware.JWT([]byte(constants.SecretKey)))
	e.PUT("homestays/update", c.UpdateHomestay(), middleware.JWT([]byte(constants.SecretKey)))
	e.POST("homestays/create", c.CreateHomestay(), middleware.JWT([]byte(constants.SecretKey)))
	e.DELETE("/homestays/delete", c.DeleteHomestay(), middleware.JWT([]byte(constants.SecretKey)))
}
