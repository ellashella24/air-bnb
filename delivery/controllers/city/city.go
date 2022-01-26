package city

import (
	"air-bnb/delivery/common"
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
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		formatCity := CityFormatResponse{}
		formatCities := []CityFormatResponse{}

		for i := 0; i < len(res); i++ {
			formatCity.ID = res[i].ID
			formatCity.Name = res[i].Name
			formatCities = append(formatCities, formatCity)
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(formatCities))
	}
}

func (uc *CityController) GetCityByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		cityID, _ := strconv.Atoi(c.Param("id"))

		res, err := uc.cityRepo.GetCityByID(cityID)

		if err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		formatCity := CityFormatResponse{
			ID:   res.ID,
			Name: res.Name,
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(formatCity))
	}
}

func (uc *CityController) CreateCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		cityRegisterReq := CreateCityRequestFormat{}

		c.Bind(&cityRegisterReq)

		err := c.Validate(&cityRegisterReq)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newCity := entities.City{}
		newCity.Name = cityRegisterReq.Name

		res, err := uc.cityRepo.CreateCity(newCity)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		formatCity := CityFormatResponse{
			ID:   res.ID,
			Name: res.Name,
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(formatCity))
	}
}

func (uc *CityController) UpdateCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		cityID, _ := strconv.Atoi(c.Param("id"))

		cityUpdateReq := UpdateCityRequestFormat{}

		c.Bind(&cityUpdateReq)

		err := c.Validate(&cityUpdateReq)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updatedCity := entities.City{}
		updatedCity.Name = cityUpdateReq.Name

		res, err := uc.cityRepo.UpdateCity(cityID, updatedCity)

		if err != nil || res.Name == "" {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		formatCity := CityFormatResponse{
			ID:   uint(cityID),
			Name: res.Name,
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(formatCity))
	}
}

func (uc *CityController) DeleteCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		cityID, _ := strconv.Atoi(c.Param("id"))

		_, err := uc.cityRepo.DeleteCity(cityID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}
