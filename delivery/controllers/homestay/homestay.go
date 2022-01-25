package homestay

import (
	"air-bnb/delivery/middlewares"
	"air-bnb/entities"
	"air-bnb/repository/homestay"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
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
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gagal gaes get all homestay",
				"data":    err,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all homestay",
			"data":    res,
		})
	}
}

func (s StructCtrlHomestay) GetHostHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userId := ExtractTokenUserId(c)
		userId := middlewares.NewAuth().ExtractTokenUserID(c)
		res, err := s.repository.GetHostHomestay(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gagal gaes get all host homestay",
				"data":    err,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all host homestay",
			"data":    res,
		})
	}

}

func (s StructCtrlHomestay) CreateHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := middlewares.NewAuth().ExtractTokenUserID(c)

		var homestay FormReqCreate
		err := c.Bind(&homestay)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "salah input gaess",
			})
		}
		var create entities.Homestay
		create.Name = homestay.Name
		create.Price = homestay.Price
		create.City_id = homestay.CityId
		create.HostID = uint(userId)

		res, err2 := s.repository.CreteaHomestay(create)
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gagal create homestay gas",
				"data":    err2,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "ok mantap create success",
			"data":    res,
		})
	}
}

func (s StructCtrlHomestay) UpdateHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userId := ExtractTokenUserId(c)
		userId := middlewares.NewAuth().ExtractTokenUserID(c)

		fond, err := s.repository.GetHomestayId(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gagal gaes menemukan id homestay",
			})
		}

		update := FormReqUpdate{}
		err2 := c.Bind(&update)
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "salah input gaess",
			})
		}
		fond.Name = update.Name
		fond.Price = update.Price

		res, err := s.repository.UpdateHomestay(fond)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "gagal save update gas",
				"data":    err,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "ok mantap update success",
			"data":    res,
		})
	}

}

func (s StructCtrlHomestay) DeleteHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := middlewares.NewAuth().ExtractTokenUserID(c)

		err := s.repository.DeleteHomestay(userId)
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

//--------------------------------------------------------------------------------
//TOKEN
//--------------------------------------------------------------------------------

func ExtractTokenUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	// fmt.Println(user)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		fmt.Println(claims)
		return userId
	}
	return -1
}
