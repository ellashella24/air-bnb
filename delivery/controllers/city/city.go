package city

import (
	"air-bnb/entities"
	"air-bnb/repository/city"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CityController struct {
	cityRepo city.City
}

func NewCityController(cityRepo city.City) *CityController {
	return &CityController{cityRepo}
}

func (uc *CityController) GetAllCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.cityRepo.GetAllCity()

		if err != nil || len(res) == 0 {
			c.JSON(http.StatusBadRequest, "error")
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *CityController) GetCityByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		cityID, _ := strconv.Atoi(c.Param("id"))

		res, err := uc.cityRepo.GetCityByID(cityID)

		if err != nil || res.ID == 0 {
			c.JSON(http.StatusBadRequest, "error")
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *CityController) CreateCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		cityRegisterReq := CreateCityRequestFormat{}

		c.Bind(&cityRegisterReq)

		err := c.Validate(&cityRegisterReq)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "error")
		}

		newCity := entities.City{}
		newCity.Name = cityRegisterReq.Name

		res, err := uc.cityRepo.CreateCity(newCity)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "error")
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *CityController) UpdateCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		cityID, _ := strconv.Atoi(c.Param("id"))

		cityUpdateReq := UpdateCityRequestFormat{}

		c.Bind(&cityUpdateReq)

		err := c.Validate(&cityUpdateReq)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "error")
		}

		updatedCity := entities.City{}
		updatedCity.Name = cityUpdateReq.Name

		res, err := uc.cityRepo.UpdateCity(cityID, updatedCity)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "can't update city")
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *CityController) DeleteCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		cityID, _ := strconv.Atoi(c.Param("id"))

		_, err := uc.cityRepo.DeleteCity(cityID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "can't delete city")
		}

		return c.JSON(http.StatusOK, "success delete city")
	}
}
