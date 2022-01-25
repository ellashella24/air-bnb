package routes

import (
	controllers "air-bnb/delivery/controllers/homestay"

	"github.com/labstack/echo/v4"
)

func PathHomestay(e *echo.Echo, c controllers.StructCtrlHomestay) {

	e.GET("/homestays", c.GetAllHomestay())
	e.GET("/homestays/host", c.GetHostHomestay())
	e.PUT("homestays/update", c.UpdateHomestay())
	e.POST("homestays/create", c.CreateHomestay())
	e.DELETE("/homestays/delete", c.DeleteHomestay())
}
