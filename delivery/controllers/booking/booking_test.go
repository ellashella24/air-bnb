package booking

import (
	"air-bnb/configs"
	"air-bnb/entities"
	"air-bnb/utils"
	"gorm.io/gorm"
	"testing"
)

var jwtToken = ""

func TestSetup(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.City{})
	db.Migrator().DropTable(&entities.Homestay{})
	db.Migrator().DropTable(&entities.Booking{})

	db.AutoMigrate(entities.User{})
	db.AutoMigrate(entities.City{})
	db.AutoMigrate(entities.Homestay{})
	db.AutoMigrate(entities.Booking{})

	var user entities.User
	user = entities.User{
		Model: gorm.Model{},
		Name:  "naufal",
		Email: "naufal@gmail.com",
		Role:  "user",
	}

	var city entities.City
	city = entities.City{
		Name:       "indramayu",
		HomestayID: nil,
	}

	var homestay entities.Homestay
	homestay = entities.Homestay{
		Name:           "koskosan",
		Price:          100000,
		Booking_Status: "available",
		HostID:         1,
		City_id:        1,
	}

	db.Create(&user)
	db.Create(&city)
	db.Create(&homestay)

}

func TestCreate(t *testing.T) {

}