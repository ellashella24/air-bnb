package user

import (
	"air-bnb/constants"
	"air-bnb/delivery/common"
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

func TestUserController(t *testing.T) {
	t.Run("1. Success Register Test", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(RegisterFormatRequest{
			Name:     "user1",
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/register")

		userController := NewUserController(mockUserRepository{})
		userController.Register()(context)

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("2. Error Register Test 1 - Error Validate ", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(RegisterFormatRequest{
			Name:     "user1",
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/register")

		userController := NewUserController(mockFalseUserRepository{})
		userController.Register()(context)

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("3. Error Register Test 2 - Can't Register ", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(RegisterFormatRequest{
			Name:     "user1",
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/register")

		userController := NewUserController(mockFalseUserRepository{})
		userController.Register()(context)

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	jwtToken := ""
	t.Run("4. Success Login Test", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(LoginFormatRequest{
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userController := NewUserController(mockUserRepository{})
		userController.Login()(context)

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := (response.Data).(map[string]interface{})

		jwtToken = data["token"].(string)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("5. Error Login Test 1 - Error Validate", func(t *testing.T) {
		e := echo.New()
		// e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(LoginFormatRequest{
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userController := NewUserController(mockFalseUserRepository{})
		userController.Login()(context)

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("6. Error Login Test 2 - Not Found", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(LoginFormatRequest{
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userController := NewUserController(mockFalseUserRepository{})
		userController.Login()(context)

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("7. Error Login Test 3 - Password doesn't match", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(LoginFormatRequest{
			Email:    "user1@mail.com",
			Password: "ok",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userController := NewUserController(mockUserRepository{})
		userController.Login()(context)

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		fmt.Println(response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("7. Success Get All User Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockUserRepository{})
		mw := middlewares.NewAuth()

		err := (mw.IsAdmin)(userController.GetAllUser())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("8. Error Get All User Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockFalseUserRepository{})
		mw := middlewares.NewAuth()

		err := (mw.IsAdmin)(userController.GetAllUser())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("9. Success Get User Profile", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockUserRepository{})

		err := middleware.JWT([]byte(constants.SecretKey))(userController.GetUserByID())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("10. Error Get User Profile Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockFalseUserRepository{})

		err := middleware.JWT([]byte(constants.SecretKey))(userController.GetUserByID())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("11. Success Update User Test", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(UpdateFormatRequest{
			Name:     "user1 new",
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockUserRepository{})

		err := middleware.JWT([]byte(constants.SecretKey))(userController.UpdateUser())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("12. Error Update User Test 1 - Error Validate ", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(UpdateFormatRequest{
			Name:     "user1 new",
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockFalseUserRepository{})

		err := middleware.JWT([]byte(constants.SecretKey))(userController.UpdateUser())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("13. Error Update User Test 2 - Can't Update", func(t *testing.T) {
		e := echo.New()
		e.Validator = &CustomValidator{validator: validator.New()}

		requestBody, _ := json.Marshal(UpdateFormatRequest{
			Name:     "user1 new",
			Email:    "user1@mail.com",
			Password: "user1",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockFalseUserRepository{})

		err := middleware.JWT([]byte(constants.SecretKey))(userController.UpdateUser())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.DefaultResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("14. Success Delete User Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockUserRepository{})

		err := middleware.JWT([]byte(constants.SecretKey))(userController.DeleteUser())(context)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response common.ResponseSuccess

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("15. Error Delete User Test", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockFalseUserRepository{})

		err := middleware.JWT([]byte(constants.SecretKey))(userController.DeleteUser())(context)

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

type mockFalseUserRepository struct{}

func (mfu mockFalseUserRepository) GetAllUser() ([]entities.User, error) {
	return nil, errors.New("can't get users data")
}

func (mfu mockFalseUserRepository) GetUserByID(userID int) (entities.User, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return entities.User{
		ID: 1, Name: "user1", Email: "user1@mail.com", Password: string(hashPassword), Role: "admin"}, errors.New("can't get user data")
}

func (mfu mockFalseUserRepository) GetUserByEmail(email string) (entities.User, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return entities.User{
		ID: 1, Name: "user1", Email: "user1@mail.com", Password: string(hashPassword), Role: "admin"}, errors.New("can't get user data")
}

func (mfu mockFalseUserRepository) CreateUser(newUser entities.User) (entities.User, error) {
	// hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return entities.User{
		ID: 0, Name: "", Email: "", Password: "", Role: ""}, errors.New("can't create user data")
}

func (mfu mockFalseUserRepository) UpdateUser(userID int, updatedUser entities.User) (entities.User, error) {
	// hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return entities.User{
		ID: 0, Name: "", Email: "", Password: "", Role: ""}, errors.New("can't update user data")
}

func (mfu mockFalseUserRepository) DeleteUser(userID int) (entities.User, error) {
	// hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

	return entities.User{
		ID: 0, Name: "", Email: "", Password: "", Role: ""}, errors.New("can't delete user data")
}
