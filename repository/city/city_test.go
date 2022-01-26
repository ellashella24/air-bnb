package city

import (
	"air-bnb/configs"
	"air-bnb/entities"
	"air-bnb/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCity(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.City{})
	db.AutoMigrate(&entities.City{})

	cityRepo := NewCityRepository(db)

	t.Run("success case", func(t *testing.T) {
		mockCity := entities.City{Name: "city1"}

		_, _ = cityRepo.CreateCity(mockCity)

		res, _ := cityRepo.GetAllCity()

		assert.Equal(t, mockCity.Name, res[0].Name)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.City{})

		mockCity := entities.City{Name: "city1"}

		_, _ = cityRepo.CreateCity(mockCity)

		res, _ := cityRepo.GetAllCity()

		assert.Equal(t, []entities.City{}, res)
	})
}

func TestGetCityByID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.City{})
	db.AutoMigrate(&entities.City{})

	cityRepo := NewCityRepository(db)

	t.Run("success case", func(t *testing.T) {
		mockCity := entities.City{Name: "city1"}

		createCity, _ := cityRepo.CreateCity(mockCity)

		res, _ := cityRepo.GetCityByID(int(createCity.ID))

		assert.Equal(t, mockCity.Name, res.Name)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.City{})

		mockCity := entities.City{Name: "city1"}

		createCity, _ := cityRepo.CreateCity(mockCity)

		res, _ := cityRepo.GetCityByID(int(createCity.ID))

		assert.Equal(t, "", res.Name)
	})
}

func TestCreateCity(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.City{})
	db.AutoMigrate(&entities.City{})

	cityRepo := NewCityRepository(db)

	t.Run("success case", func(t *testing.T) {
		mockCity := entities.City{Name: "city1"}

		res, _ := cityRepo.CreateCity(mockCity)

		assert.Equal(t, mockCity.Name, res.Name)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.City{})

		mockCity := entities.City{Name: "city1"}

		res, _ := cityRepo.CreateCity(mockCity)

		assert.Equal(t, 0, int(res.ID))
	})
}

func TestUpdateCity(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.City{})
	db.AutoMigrate(&entities.City{})

	cityRepo := NewCityRepository(db)

	t.Run("success case", func(t *testing.T) {
		mockCity := entities.City{Name: "city1"}
		mockUpdateCity := entities.City{Name: "city1 new"}

		createCity, _ := cityRepo.CreateCity(mockCity)
		res, _ := cityRepo.UpdateCity(int(createCity.ID), mockUpdateCity)

		assert.Equal(t, mockUpdateCity.Name, res.Name)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.City{})

		mockCity := entities.City{Name: "city1"}
		mockUpdateCity := entities.City{Name: "city1 new"}

		createCity, _ := cityRepo.CreateCity(mockCity)
		res, _ := cityRepo.UpdateCity(int(createCity.ID), mockUpdateCity)

		assert.Equal(t, "", res.Name)
	})
}

func TestDeleteCity(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.City{})
	db.AutoMigrate(&entities.City{})

	cityRepo := NewCityRepository(db)

	t.Run("success case", func(t *testing.T) {
		mockCity := entities.City{Name: "city1"}

		createCity, _ := cityRepo.CreateCity(mockCity)

		res, err := cityRepo.DeleteCity(int(createCity.ID))

		assert.Equal(t, "", res.Name)
		assert.Equal(t, nil, err)
	})

	t.Run("error case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.City{})

		mockCity := entities.City{Name: "city1"}

		createCity, _ := cityRepo.CreateCity(mockCity)

		res, _ := cityRepo.DeleteCity(int(createCity.ID))

		assert.Equal(t, "", res.Name)
	})
}
