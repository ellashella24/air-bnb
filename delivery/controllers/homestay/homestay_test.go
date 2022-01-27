package homestay

import (
	"air-bnb/constants"
	"air-bnb/delivery/common"
	"air-bnb/delivery/controllers/user"
	"air-bnb/entities"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

var jwtToken string

func TestHomestay(t *testing.T) {
	t.Run("4. Success Login Test", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(user.LoginFormatRequest{
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userController := user.NewUserController(mockUserRepository{})
		userController.Login()(context)

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		fmt.Println(response)

		data := (response.Data).(map[string]interface{})

		jwtToken = data["token"].(string)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("getallHomestay", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/homestays")

		homestayController := NewControllerHomestay(mockHomestayRepository{})
		homestayController.GetAllHomestay()(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("getallHomestayHost", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/homestays/host")

		homestayController := NewControllerHomestay(mockHomestayRepository{})
		middleware.JWT([]byte(constants.SecretKey))(homestayController.GetAllHostHomestay())(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("createHomestay", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":   "homestay1",
			"price":  1000,
			"cityid": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/homestays/create")

		homestayController := NewControllerHomestay(mockHomestayRepository{})
		middleware.JWT([]byte(constants.SecretKey))(homestayController.CreateHomestay())(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("updateHomestay", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":   "homestay1update",
			"price":  1000,
			"cityid": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/homestays/update")

		homestayController := NewControllerHomestay(mockHomestayRepository{})
		middleware.JWT([]byte(constants.SecretKey))(homestayController.UpdateHomestay())(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("deleteHomestay", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/homestays/delete")

		homestayController := NewControllerHomestay(mockHomestayRepository{})
		middleware.JWT([]byte(constants.SecretKey))(homestayController.DeleteHomestay())(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("getHomestayByCityId", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/homestays/:search")
		context.SetParamNames("search")
		context.SetParamValues("1")

		homestayController := NewControllerHomestay(mockHomestayRepository{})
		homestayController.GetHomestayByCityId()(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}

func TestHomestayFalse(t *testing.T) {
	t.Run("4. Success Login Test", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(user.LoginFormatRequest{
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userController := user.NewUserController(mockUserRepository{})
		userController.Login()(context)

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		fmt.Println(response)

		data := (response.Data).(map[string]interface{})

		jwtToken = data["token"].(string)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("getallHomestayFalse", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/homestays")

		homestayController := NewControllerHomestay(falsemockHomestayRepository{})
		homestayController.GetAllHomestay()(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("getallHomestayHost", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/homestays/host")

		homestayController := NewControllerHomestay(falsemockHomestayRepository{})
		middleware.JWT([]byte(constants.SecretKey))(homestayController.GetAllHostHomestay())(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Not Found", responses.Message)
	})

	t.Run("createHomestay", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":   "homestay1",
			"price":  1000,
			"cityid": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/homestays/create")

		homestayController := NewControllerHomestay(mockHomestayRepository{})
		middleware.JWT([]byte(constants.SecretKey))(homestayController.CreateHomestay())(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("updateHomestay", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":   "homestay1update",
			"price":  1000,
			"cityid": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/homestays/update")

		homestayController := NewControllerHomestay(mockHomestayRepository{})
		middleware.JWT([]byte(constants.SecretKey))(homestayController.UpdateHomestay())(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("deleteHomestay", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/homestays/delete")

		homestayController := NewControllerHomestay(mockHomestayRepository{})
		middleware.JWT([]byte(constants.SecretKey))(homestayController.DeleteHomestay())(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Successful Operation", responses.Message)
	})

	t.Run("getHomestayByCityId", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/homestays/:search")
		context.SetParamNames("search")
		context.SetParamValues("1")

		homestayController := NewControllerHomestay(mockHomestayRepository{})
		homestayController.GetHomestayByCityId()(context)

		var responses common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
}

type mockHomestayRepository struct{}

func (m mockHomestayRepository) GetallHomestay() ([]entities.Homestay, error) {
	return []entities.Homestay{
		{Name: "homestay1", Price: 1000},
	}, nil
}
func (m mockHomestayRepository) GetAllHostHomestay(id int) ([]entities.Homestay, error) {
	return []entities.Homestay{
		{HostID: uint(id), Name: "homestay1", Price: 2000},
	}, nil
}

func (m mockHomestayRepository) CreteaHomestay(homestay entities.Homestay) (entities.Homestay, error) {
	return entities.Homestay{Name: "homestay1", Price: 1000, City_id: 1}, nil
}

func (m mockHomestayRepository) GetHomestayIdByHostId(idUser, idHomestay int) (entities.Homestay, error) {
	return entities.Homestay{
		Name:  "test1",
		Price: 2000,
	}, nil
}

func (m mockHomestayRepository) UpdateHomestay(homestay entities.Homestay) (entities.Homestay, error) {
	return entities.Homestay{Name: "homestay1", Price: 1000, City_id: 1}, nil
}
func (m mockHomestayRepository) DeleteHomestayByHostId(id, homeStay int) error {
	return nil
}
func (m mockHomestayRepository) GetHomestayByCityId(city string) ([]entities.Homestay, error) {
	return []entities.Homestay{
		{Name: "homestay1", Price: 2000, City_id: 1},
	}, nil
}

type falsemockHomestayRepository struct{}

func (m falsemockHomestayRepository) GetallHomestay() ([]entities.Homestay, error) {
	return []entities.Homestay{}, errors.New("tidak menemukan homstay")
}
func (m falsemockHomestayRepository) GetAllHostHomestay(id int) ([]entities.Homestay, error) {
	return []entities.Homestay{}, errors.New("tidak menemukan host homstay")
}
func (m falsemockHomestayRepository) CreteaHomestay(homestay entities.Homestay) (entities.Homestay, error) {

	return entities.Homestay{Name: "homestay1", Price: 1000}, nil
}
func (m falsemockHomestayRepository) GetHomestayIdByHostId(idUser, idHomestay int) (entities.Homestay, error) {
	return entities.Homestay{
		Name:  "test1",
		Price: 2000,
	}, nil
}
func (m falsemockHomestayRepository) UpdateHomestay(homestay entities.Homestay) (entities.Homestay, error) {
	return entities.Homestay{Name: "homestay1", Price: 1000}, nil
}
func (m falsemockHomestayRepository) DeleteHomestayByHostId(id, homeStay int) error {
	return nil
}
func (m falsemockHomestayRepository) GetHomestayByCityId(city string) ([]entities.Homestay, error) {
	return []entities.Homestay{
		{Name: "homestay1", Price: 2000},
	}, nil
}

//////MOCK USER
type mockUserRepository struct{}

func (mu mockUserRepository) GetAllUser() ([]entities.User, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return []entities.User{
		{ID: 1, Name: "user1", Email: "user1@mail.com", Password: string(hashPassword), Role: "admin"}}, nil
}

func (mu mockUserRepository) GetUserByID(userID int) (entities.User, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return entities.User{
		ID: 1, Name: "user1", Email: "user1@mail.com", Password: string(hashPassword), Role: "admin"}, nil
}

func (mu mockUserRepository) GetUserByEmail(email string) (entities.User, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return entities.User{
		ID: 1, Name: "user1", Email: "user1@mail.com", Password: string(hashPassword), Role: "admin"}, nil
}

func (mu mockUserRepository) CreateUser(newUser entities.User) (entities.User, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return entities.User{
		ID: 1, Name: "user1", Email: "user1@mail.com", Password: string(hashPassword), Role: "admin"}, nil
}

func (mu mockUserRepository) UpdateUser(userID int, updatedUser entities.User) (entities.User, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return entities.User{
		ID: 1, Name: "user1 new", Email: "user1@mail.com", Password: string(hashPassword), Role: "admin"}, nil
}

func (mu mockUserRepository) DeleteUser(userID int) (entities.User, error) {
	// hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return entities.User{
		ID: 0, Name: "", Email: "", Password: "", Role: ""}, nil
}
