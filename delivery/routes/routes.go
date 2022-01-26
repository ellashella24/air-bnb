package routes

import (
	"air-bnb/delivery/controllers/city"
	"air-bnb/delivery/controllers/user"
	"air-bnb/delivery/middlewares"

	"air-bnb/constants"
	controllers "air-bnb/delivery/controllers/homestay"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func RegisterPath(e *echo.Echo, uc *user.UserController, cc *city.CityController) {
	e.Validator = &CustomValidator{validator: validator.New()}

	mw := middlewares.NewAuth()

	e.GET("/users", uc.GetAllUser(), middleware.JWT([]byte("secret123")), mw.IsAdmin)
	e.GET("/users/:id", uc.GetUserByID(), middleware.JWT([]byte("secret123")))
	e.POST("/users/register", uc.Register())
	e.POST("/users/login", uc.Login())
	e.PUT("/users", uc.UpdateUser(), middleware.JWT([]byte("secret123")))
	e.DELETE("/users/:id", uc.DeleteUser(), middleware.JWT([]byte("secret123")))

	e.GET("/city", cc.GetAllCity())
	e.GET("/city/:id", cc.GetCityByID())
	e.POST("/city", cc.CreateCity(), middleware.JWT([]byte("secret123")), mw.IsAdmin)
	e.PUT("/city/:id", cc.UpdateCity(), middleware.JWT([]byte("secret123")), mw.IsAdmin)
	e.DELETE("/city/:id", cc.DeleteCity(), middleware.JWT([]byte("secret123")), mw.IsAdmin)

}

func PathHomestay(e *echo.Echo, c controllers.StructCtrlHomestay) {

	e.GET("/homestays", c.GetAllHomestay())
	e.GET("/homestays/host", c.GetAllHostHomestay(), middleware.JWT([]byte(constants.SecretKey)))
	e.PUT("homestays/update", c.UpdateHomestay(), middleware.JWT([]byte(constants.SecretKey)))
	e.POST("homestays/create", c.CreateHomestay(), middleware.JWT([]byte(constants.SecretKey)))
	e.DELETE("/homestays/delete", c.DeleteHomestay(), middleware.JWT([]byte(constants.SecretKey)))
	e.GET("/homestays/", c.GetHomestayByCityId())
}
