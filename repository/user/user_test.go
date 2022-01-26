package user

import (
	"air-bnb/configs"
	"air-bnb/entities"
	"air-bnb/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGetAllUser(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	userRepo := NewUserRepository(db)

	t.Run("success case", func(t *testing.T) {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}

		_, _ = userRepo.CreateUser(mockUser)

		res, _ := userRepo.GetAllUser()

		assert.Equal(t, mockUser.Name, res[0].Name)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}

		_, _ = userRepo.CreateUser(mockUser)

		res, _ := userRepo.GetAllUser()

		assert.Equal(t, []entities.User{}, res)
	})
}

func TestGetUserByID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	userRepo := NewUserRepository(db)

	t.Run("success case", func(t *testing.T) {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}

		createUser, _ := userRepo.CreateUser(mockUser)

		res, _ := userRepo.GetUserByID(int(createUser.ID))

		assert.Equal(t, mockUser.Name, res.Name)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}

		createUser, _ := userRepo.CreateUser(mockUser)

		res, _ := userRepo.GetUserByID(int(createUser.ID))

		assert.Equal(t, "", res.Name)
	})
}

func TestGetUserByEmail(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	userRepo := NewUserRepository(db)

	t.Run("success case", func(t *testing.T) {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}

		createUser, _ := userRepo.CreateUser(mockUser)

		res, _ := userRepo.GetUserByEmail(createUser.Email)

		assert.Equal(t, mockUser.Name, res.Name)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}

		createUser, _ := userRepo.CreateUser(mockUser)

		res, _ := userRepo.GetUserByEmail(createUser.Email)

		assert.Equal(t, entities.User{}, res)
	})
}

func TestCreateUser(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	userRepo := NewUserRepository(db)

	t.Run("success case", func(t *testing.T) {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}

		res, _ := userRepo.CreateUser(mockUser)

		assert.Equal(t, mockUser.Name, res.Name)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}

		res, _ := userRepo.CreateUser(mockUser)

		assert.Equal(t, 0, int(res.ID))
	})
}

func TestUpdateUser(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	userRepo := NewUserRepository(db)

	t.Run("success case", func(t *testing.T) {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}
		mockUpdateUser := entities.User{Name: "user1 new", Email: "user1@mail.com", Password: string(hashPassword)}

		createUser, _ := userRepo.CreateUser(mockUser)
		res, _ := userRepo.UpdateUser(int(createUser.ID), mockUpdateUser)

		assert.Equal(t, mockUpdateUser.Name, res.Name)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}
		mockUpdateUser := entities.User{Name: "user1 new", Email: "user1@mail.com", Password: string(hashPassword)}

		createUser, _ := userRepo.CreateUser(mockUser)
		res, _ := userRepo.UpdateUser(int(createUser.ID), mockUpdateUser)

		assert.Equal(t, "", res.Name)
	})
}

func TestDeleteUser(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	userRepo := NewUserRepository(db)

	t.Run("success case", func(t *testing.T) {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}

		createUser, _ := userRepo.CreateUser(mockUser)

		res, err := userRepo.DeleteUser(int(createUser.ID))

		assert.Equal(t, "", res.Name)
		assert.Equal(t, nil, err)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.User{})

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte("user1"), bcrypt.MinCost)

		mockUser := entities.User{Name: "user1", Email: "user1@mail.com", Password: string(hashPassword)}

		createUser, _ := userRepo.CreateUser(mockUser)

		res, _ := userRepo.DeleteUser(int(createUser.ID))

		assert.Equal(t, "", res.Name)
	})
}
