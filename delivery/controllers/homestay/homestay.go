package homestay

import (
	"air-bnb/delivery/common"
	"air-bnb/delivery/middlewares"
	"air-bnb/entities"
	"air-bnb/repository/homestay"
	"net/http"
	"strconv"

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
		create.CityID = homestay.CityId
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
		homestayId, _ := strconv.Atoi(c.Param("id"))
		fond, err := s.repository.GetHomestayIdByHostId(userId, homestayId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(400, "host ID not found"))
		}
		update := FormReqUpdate{}
		err2 := c.Bind(&update)
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(400, "wrong input"))
		}
		fond.Name = update.Name
		fond.Price = update.Price
		fond.City_id = update.CityId
		res, err := s.repository.UpdateHomestay(fond)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, common.SuccessResponse(res))
	}

}

func (s StructCtrlHomestay) DeleteHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := middlewares.NewAuth().ExtractTokenUserID(c)
		homestayId, _ := strconv.Atoi(c.Param("id"))
		err := s.repository.DeleteHomestayByHostId(userId, homestayId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(400, "host/homestay ID tidak ditemukan"))
		}
		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}

func (s StructCtrlHomestay) GetHomestayByCityId() echo.HandlerFunc {
	return func(c echo.Context) error {
		search := c.QueryParam("search")

		res, err := s.repository.GetHomestayByCityId(search)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(400, "kota tidak ditemukan"))
		}
		return c.JSON(http.StatusOK, common.SuccessResponse(res))

	}
}
