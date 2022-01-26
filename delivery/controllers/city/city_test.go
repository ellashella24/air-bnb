package city

import (
	"air-bnb/delivery/common"
	"air-bnb/delivery/controllers/user"
	"air-bnb/delivery/middlewares"
	"air-bnb/entities"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func TestCityController(t *testing.T) {
	t.Run("1. Success Get All City Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/city")

		cityController := NewCityController(mockCityRepository{})
		cityController.GetAllCity()(context)

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("2. Error Get All User Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/city")

		cityController := NewCityController(mockFalseCityRepository{})
		cityController.GetAllCity()(context)

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("3. Success Get City Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/city")
		context.SetParamNames("id")
		context.SetParamValues("1")

		cityController := NewCityController(mockCityRepository{})
		cityController.GetCityByID()(context)

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("4. Error Get City Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/city")
		context.SetParamNames("id")
		context.SetParamValues("1")

		cityController := NewCityController(mockFalseCityRepository{})
		cityController.GetCityByID()(context)

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Not Found", response.Message)
	})
	jwtToken := ""
	t.Run("5. Success Login Test", func(t *testing.T) {
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

		data := (response.Data).(map[string]interface{})

		jwtToken = data["token"].(string)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("6. Success Create City Test", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(CreateCityRequestFormat{
			Name: "city1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/city")

		userController := NewCityController(mockCityRepository{})
		mw := middlewares.NewAuth()

		err := (mw.IsAdmin)(userController.CreateCity())(context)

		if err != nil {
			fmt.Println(err)
			return
		}
		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("7. Error Create City Test 1 - Can't create city", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(CreateCityRequestFormat{
			Name: "city1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewCityController(mockFalseCityRepository{})
		mw := middlewares.NewAuth()

		err := (mw.IsAdmin)(userController.CreateCity())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("8. Error Create City Test 2 - Error Validate", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(CreateCityRequestFormat{
			Name: "city1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewCityController(mockFalseCityRepository{})
		mw := middlewares.NewAuth()

		err := (mw.IsAdmin)(userController.CreateCity())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("9. Success Update City Test", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(CreateCityRequestFormat{
			Name: "city1 new",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/city")

		userController := NewCityController(mockCityRepository{})
		mw := middlewares.NewAuth()

		err := (mw.IsAdmin)(userController.UpdateCity())(context)

		if err != nil {
			fmt.Println(err)
			return
		}
		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("7. Error Update City Test 1 - Can't update city", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(CreateCityRequestFormat{
			Name: "city1 new",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewCityController(mockFalseCityRepository{})
		mw := middlewares.NewAuth()

		err := (mw.IsAdmin)(userController.UpdateCity())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("8. Error Update City Test 2 - Error Validate", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(CreateCityRequestFormat{
			Name: "city1",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewCityController(mockFalseCityRepository{})
		mw := middlewares.NewAuth()

		err := (mw.IsAdmin)(userController.UpdateCity())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("9. Success Delete City Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/city")
		context.SetParamNames("id")
		context.SetParamValues("1")

		cityController := NewCityController(mockCityRepository{})
		mw := middlewares.NewAuth()

		err := (mw.IsAdmin)(cityController.DeleteCity())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("10. Error Delete City Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/city")
		context.SetParamNames("id")
		context.SetParamValues("1")

		cityController := NewCityController(mockFalseCityRepository{})
		mw := middlewares.NewAuth()

		err := (mw.IsAdmin)(cityController.DeleteCity())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
}

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

type mockCityRepository struct{}

func (mc mockCityRepository) GetAllCity() ([]entities.City, error) {
	return []entities.City{
		{ID: 1, Name: "city1"}}, nil
}

func (mc mockCityRepository) GetCityByID(cityID int) (entities.City, error) {
	return entities.City{
		ID: 1, Name: "city1"}, nil
}

func (mu mockCityRepository) CreateCity(newCity entities.City) (entities.City, error) {
	return entities.City{
		ID: 1, Name: "city1"}, nil
}

func (mu mockCityRepository) UpdateCity(cityID int, updatedCity entities.City) (entities.City, error) {
	return entities.City{
		ID: 1, Name: "city1 new"}, nil
}

func (mu mockCityRepository) DeleteCity(userID int) (entities.City, error) {
	return entities.City{
		ID: 0, Name: ""}, nil
}

type mockFalseCityRepository struct{}

func (mfc mockFalseCityRepository) GetAllCity() ([]entities.City, error) {
	return []entities.City{
		{ID: 0, Name: ""}}, errors.New("can't get cities data")
}

func (mfc mockFalseCityRepository) GetCityByID(cityID int) (entities.City, error) {
	return entities.City{
		ID: 0, Name: ""}, errors.New("can't get city data")
}

func (mfu mockFalseCityRepository) CreateCity(newCity entities.City) (entities.City, error) {
	return entities.City{
		ID: 0, Name: ""}, errors.New("can't create city data")
}

func (mfu mockFalseCityRepository) UpdateCity(cityID int, updatedCity entities.City) (entities.City, error) {
	return entities.City{
		ID: 0, Name: ""}, errors.New("can't update city data")
}

func (mfu mockFalseCityRepository) DeleteCity(userID int) (entities.City, error) {
	return entities.City{
		ID: 0, Name: ""}, errors.New("can't get city data")
}
