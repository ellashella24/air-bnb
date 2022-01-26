package city

import (
	"air-bnb/entities"

	"gorm.io/gorm"
)

type City interface {
	GetAllCity() ([]entities.City, error)
	GetCityByID(cityID int) (entities.City, error)
	CreateCity(city entities.City) (entities.City, error)
	UpdateCity(cityID int, city entities.City) (entities.City, error)
	DeleteCity(cityID int) (entities.City, error)
}

type cityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *cityRepository {
	return &cityRepository{db}
}

func (ur *cityRepository) GetAllCity() ([]entities.City, error) {
	cities := []entities.City{}

	err := ur.db.Find(&cities).Error

	if err != nil {
		return cities, nil
	}

	return cities, nil
}

func (ur *cityRepository) GetCityByID(cityID int) (entities.City, error) {
	city := entities.City{}

	err := ur.db.Where("ID = ?", cityID).Find(&city).Error

	if err != nil {
		return city, nil
	}

	return city, nil
}

func (ur *cityRepository) CreateCity(newCity entities.City) (entities.City, error) {
	ur.db.Save(&newCity)

	return newCity, nil
}

func (ur *cityRepository) UpdateCity(cityID int, updatedCity entities.City) (entities.City, error) {
	city := entities.City{}

	err := ur.db.Where("id = ?", cityID).Find(&city).Error

	if err != nil || city.ID == 0 {
		return city, err
	}

	ur.db.Model(&city).Updates(updatedCity)

	return updatedCity, nil
}

func (ur *cityRepository) DeleteCity(cityID int) (entities.City, error) {
	city := entities.City{}

	err := ur.db.Where("id = ?", cityID).Delete(&city).Error

	if err != nil {
		return city, err
	}

	return city, nil
}
