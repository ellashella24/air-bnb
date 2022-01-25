package user

import (
	"air-bnb/delivery/middlewares"
	"air-bnb/entities"
	"air-bnb/repository/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userRepo user.User
}

func NewUserController(userRepo user.User) *UserController {
	return &UserController{userRepo}
}

func (uc *UserController) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.userRepo.GetAllUser()

		if err != nil || len(res) == 0 {
			c.JSON(http.StatusBadRequest, "error")
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.NewAuth().ExtractTokenUserID(c)

		res, err := uc.userRepo.GetUserByID(userID)

		if err != nil || res.ID == 0 {
			c.JSON(http.StatusBadRequest, "error")
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		userRegisterReq := RegisterFormatRequest{}

		c.Bind(&userRegisterReq)

		err := c.Validate(&userRegisterReq)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "error")
		}

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userRegisterReq.Password), bcrypt.MinCost)

		newUser := entities.User{}
		newUser.Name = userRegisterReq.Name
		newUser.Email = userRegisterReq.Email
		newUser.Password = string(hashPassword)

		res, err := uc.userRepo.CreateUser(newUser)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "error")
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		userLoginReq := LoginFormatRequest{}

		c.Bind(&userLoginReq)

		err := c.Validate(userLoginReq)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "error")
		}

		res, err := uc.userRepo.GetUserByEmail(userLoginReq.Email)

		if err != nil || res.ID == 0 {
			return c.JSON(http.StatusBadRequest, "not found")
		}

		err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(userLoginReq.Password))

		if err != nil {
			return c.JSON(http.StatusBadRequest, "password doesn't match")
		}

		token, _ := middlewares.NewAuth().GenerateToken(int(res.ID), res.Email, res.Role)

		return c.JSON(http.StatusOK, token)
	}
}

func (uc *UserController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.NewAuth().ExtractTokenUserID(c)

		fmt.Println(userID)

		userUpdateReq := UpdateFormatRequest{}

		c.Bind(&userUpdateReq)

		err := c.Validate(&userUpdateReq)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "error")
		}

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userUpdateReq.Password), bcrypt.MinCost)

		updatedUser := entities.User{}
		updatedUser.Name = userUpdateReq.Name
		updatedUser.Email = userUpdateReq.Email
		updatedUser.Password = string(hashPassword)

		res, err := uc.userRepo.UpdateUser(userID, updatedUser)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "can't update user")
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.NewAuth().ExtractTokenUserID(c)

		_, err := uc.userRepo.DeleteUser(userID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "can't delete user")
		}

		return c.JSON(http.StatusOK, "success delete user")
	}
}
