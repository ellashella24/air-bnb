package homestay

import (
	"air-bnb/configs"
	"air-bnb/entities"
	"air-bnb/repository/city"
	"air-bnb/repository/user"
	"air-bnb/utils"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomestay(t *testing.T) {

	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.City{})
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Homestay{})
	db.AutoMigrate(&entities.City{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Homestay{})

	cityRepo := city.NewCityRepository(db)
	userRepo := user.NewUserRepository(db)
	homestayRepo := NewRepositoryHomestay(db)

	var city entities.City
	city.Name = "Padang"
	// city.ID = 1
	db.Save(&city)

	var user entities.User
	user.Name = "Yoga"
	// user.ID = 1
	db.Save(&user)

	rp := NewRepositoryHomestay(db)

	// var dummyHomestay entities.Homestay
	// dummyHomestay.Name = "test1"
	// dummyHomestay.Price = 1000

	t.Run("CreateHomestay", func(t *testing.T) {
		var dummy entities.Homestay
		dummy.Name = "test1"
		dummy.Price = 1000
		dummy.HostID = 1
		dummy.City_id = 1

		res, err := rp.CreteaHomestay(dummy)
		assert.Nil(t, err)
		assert.Equal(t, "test1", res.Name)
	})

	t.Run("GetallHomestay", func(t *testing.T) {
		res, err := rp.GetallHomestay()
		assert.Nil(t, err)
		assert.Equal(t, "test1", res[0].Name)

	})

	t.Run("GetallHomestayHost", func(t *testing.T) {
		res, err := rp.GetAllHostHomestay(1)
		assert.Nil(t, err)
		assert.Equal(t, "test1", res[0].Name)
	})

	t.Run("GethomestayIdByHostId", func(t *testing.T) {
		res, _ := rp.GetHomestayIdByHostId(1, 1)

		assert.Equal(t, "test1", res.Name)
	})
	t.Run("UpdateHomestay", func(t *testing.T) {
		var dummy entities.Homestay
		dummy.Name = "test2"
		dummy.Price = 1000
		dummy.HostID = 1
		dummy.City_id = 1
		res, err := rp.UpdateHomestay(dummy)
		assert.Nil(t, err)
		assert.Equal(t, "test2", res.Name)
	})

	t.Run("DeleteHomestay", func(t *testing.T) {

		err := rp.DeleteHomestayByHostId(1, 1)
		assert.Nil(t, err)
	})

	t.Run("GetHomestayByCityId", func(t *testing.T) {
		mockCity := entities.City{Name: "padang"}
		createCity, _ := cityRepo.CreateCity(mockCity)

		mockUser := entities.User{Name: "User 1", Email: "user1@mail.com", Password: "user1", Role: "member"}
		createUser, _ := userRepo.CreateUser(mockUser)

		mockHomestay := entities.Homestay{Name: "Homestay 1", Price: 1000, Booking_Status: "available", HostID: createUser.ID, City_id: createCity.ID}
		_, _ = homestayRepo.CreteaHomestay(mockHomestay)

		res, _ := homestayRepo.GetHomestayByCityId(createCity.Name)

		assert.Equal(t, "test2", res[0].Name)

		// var cities
		// res, _ := rp.GetHomestayByCityId("hotel")
		// fmt.Println(res)
		// assert.Equal(t, "test1", res)
	})
}

func TestHomestayFalse(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.City{})
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Homestay{})
	db.AutoMigrate(&entities.City{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Homestay{})

	var city entities.City
	city.Name = "Padang"
	// city.ID = 1
	db.Save(&city)

	var user entities.User
	user.Name = "Yoga"
	// user.ID = 1
	db.Save(&user)

	rp := NewRepositoryHomestay(db)
	var dummyHomestay entities.Homestay
	dummyHomestay.Name = "homestay2"
	dummyHomestay.Price = 1000

	t.Run("CreateHomestayFalse", func(t *testing.T) {
		var dummy entities.Homestay
		dummy.Name = "test1"
		dummy.Price = 1000
		dummy.HostID = 1
		dummy.City_id = 1
		db.Migrator().DropTable(&entities.Homestay{})
		// db.AutoMigrate(&entities.Homestay{})
		res, err := rp.CreteaHomestay(dummy)
		assert.NotNil(t, err)
		// assert.Nil(t, res)
		assert.Equal(t, uint(0), res.ID)

	})

	t.Run("GetAllHomestayFalse", func(t *testing.T) {

		res, err := rp.GetallHomestay()
		assert.Equal(t, 0, len(res))
		assert.Equal(t, errors.New("nilai kosong"), err)

	})

	t.Run("GetAllHostHomestayFalse", func(t *testing.T) {

		res, err := rp.GetAllHostHomestay(2)
		assert.Equal(t, 0, len(res))
		assert.Equal(t, errors.New("nilai kosong"), err)

	})

	t.Run("UpdateHomestay", func(t *testing.T) {
		var update entities.Homestay
		update.Name = "homestay2"
		update.Price = 2000
		update.HostID = 1
		update.City_id = 1

		res, _ := rp.UpdateHomestay(update)
		assert.Equal(t, dummyHomestay.Name, res.Name)
		// assert.Nil(t, err)

	})

	t.Run("getHomestayHost", func(t *testing.T) {
		res, _ := rp.GetAllHostHomestay(2)
		// assert.Nil(t, err)
		assert.Equal(t, 0, len(res))
	})
	t.Run("deleteHomestayByIdFalse", func(t *testing.T) {
		err := rp.DeleteHomestayByHostId(10, 2)
		// assert.Nil(t, err)
		assert.Equal(t, errors.New("gak ketemu idnya"), err)
	})

	t.Run("getHomestayHost", func(t *testing.T) {
		res, _ := rp.GetAllHostHomestay(2)
		// assert.Nil(t, err)
		assert.Equal(t, 0, len(res))
	})

	t.Run("GethomestayIdByHostIdFalse", func(t *testing.T) {
		_, err := rp.GetHomestayIdByHostId(2, 2)

		assert.Equal(t, errors.New("gak ketemu idnya"), err)
	})

	t.Run("GetHomestayByCityIdFalse", func(t *testing.T) {
		_, err := rp.GetHomestayByCityId("goa")

		assert.Equal(t, errors.New("kota tidak ditemukan"), err)
	})
}
