package homestay

import (
	"air-bnb/delivery/common"
	"air-bnb/delivery/middlewares"
	"air-bnb/entities"
	"air-bnb/repository/homestay"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StructCtrlHomestay struct {
	repository homestay.InterfaceHomestay
}

func NewControllerHomestay(homestay homestay.InterfaceHomestay) *StructCtrlHomestay {
	return &StructCtrlHomestay{homestay}
}

func (s StructCtrlHomestay) GetAllHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := s.repository.GetallHomestay()
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewNotFoundResponse())
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(res))
	}
}

func (s StructCtrlHomestay) GetAllHostHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userId := ExtractTokenUserId(c)
		userId := middlewares.NewAuth().ExtractTokenUserID(c)
		res, err := s.repository.GetAllHostHomestay(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewNotFoundResponse())
		}
		return c.JSON(http.StatusOK, common.SuccessResponse(res))
	}

}

func (s StructCtrlHomestay) CreateHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := middlewares.NewAuth().ExtractTokenUserID(c)

		var homestay FormReqCreate
		err := c.Bind(&homestay)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(400, "wrong input"))
		}
		var create entities.Homestay
		create.Name = homestay.Name
		create.Price = homestay.Price
		create.City_id = homestay.CityId
		create.HostID = uint(userId)

		res, err2 := s.repository.CreteaHomestay(create)
		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse(500, "failed create homestay"))
		}
		return c.JSON(http.StatusOK, common.SuccessResponse(res))
	}
}

func (s StructCtrlHomestay) UpdateHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userId := ExtractTokenUserId(c)
		userId := middlewares.NewAuth().ExtractTokenUserID(c)

		fond, err := s.repository.GetHomestayIdByHostId(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(400, "host ID not found"))
		}
		fmt.Println(fond)
		update := FormReqUpdate{}
		err2 := c.Bind(&update)
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(400, "wrong input"))
		}
		fond.Name = update.Name
		fond.Price = update.Price

		res, err := s.repository.UpdateHomestay(fond)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse(500, "failed update homestay"))
		}
		return c.JSON(http.StatusOK, common.SuccessResponse(res))
	}

}

func (s StructCtrlHomestay) DeleteHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := middlewares.NewAuth().ExtractTokenUserID(c)

		err := s.repository.DeleteHomestayByHostId(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gagal delete gas",
				"data":    err,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "delete success gaess",
		})
	}
}

func (s StructCtrlHomestay) GetHomestayByCityId() echo.HandlerFunc {
	return func(c echo.Context) error {
		search := c.QueryParam("search")

		res, err := s.repository.GetHomestayByCityId(search)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ndak ada citynya",
				"data":    err,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "ada city",
			"data":    res,
		})

	}
}
