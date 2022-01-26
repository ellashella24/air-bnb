package homestay

import (
	"air-bnb/entities"
	"errors"

	"gorm.io/gorm"
)

//--------------------------------------------------------------------------------
//INTERFACE
//--------------------------------------------------------------------------------
type InterfaceHomestay interface {
	GetallHomestay() ([]entities.Homestay, error)
	GetAllHostHomestay(id int) ([]entities.Homestay, error)

	// CREATE
	CreteaHomestay(homestay entities.Homestay) (entities.Homestay, error)
	// UPDATE
	GetHomestayIdByHostId(id int) (entities.Homestay, error)
	UpdateHomestay(homestay entities.Homestay) (entities.Homestay, error)
	// DELETE
	DeleteHomestayByHostId(id int) error
	//GET HOMESTAY BY CITY
	GetHomestayByCityId(city string) ([]entities.Homestay, error)
}

type homestay struct {
	db *gorm.DB
}

func NewRepositoryHomestay(db *gorm.DB) *homestay {
	return &homestay{db}
}

func (h *homestay) GetallHomestay() ([]entities.Homestay, error) {
	var homestays []entities.Homestay
	err := h.db.Find(&homestays).Error
	if err != nil {
		return nil, errors.New("nilai kosong")
	}
	// if len(homestays) == 0 {
	// 	return nil, errors.New("data kosong")
	// }
	return homestays, nil
}

func (h *homestay) GetAllHostHomestay(id int) ([]entities.Homestay, error) {
	var homestays []entities.Homestay
	err := h.db.Where("host_id = ?", id).Find(&homestays).Error
	if err != nil {
		return nil, errors.New("nilai kosong")
	}
	return homestays, nil
}

func (h *homestay) CreteaHomestay(homestay entities.Homestay) (entities.Homestay, error) {
	err := h.db.Save(&homestay).Error
	if err != nil {
		return homestay, err
	}
	// if homestay.ID == 0 {
	// 	return homestay, errors.New("kosong")
	// }
	return homestay, nil
}

func (h *homestay) GetHomestayIdByHostId(id int) (entities.Homestay, error) {
	var homestay entities.Homestay
	err := h.db.Where("host_id = ?", id).Find(&homestay).Error
	if err != nil {
		return homestay, err
	}

	return homestay, nil
}

func (h *homestay) UpdateHomestay(homestay entities.Homestay) (entities.Homestay, error) {
	err := h.db.Save(&homestay).Error
	if err != nil {
		return homestay, err
	}
	return homestay, nil
}

func (h *homestay) DeleteHomestayByHostId(id int) error {
	var delete entities.Homestay
	err := h.db.Where("host_id = ?", id).Delete(&delete).Error
	if err != nil {
		return errors.New("gak ketemu idnya")
	}
	return nil
}

func (h *homestay) GetHomestayByCityId(city string) ([]entities.Homestay, error) {
	var homestay []entities.Homestay
	err := h.db.Table("homestays").Select("homestays.*").Joins("join cities on homestays.city_id  = cities.id").Where("cities.name = ?", city).Find(&homestay).Error
	if err != nil {
		return nil, errors.New("kota tidak ditemukan")
	}
	return homestay, nil
}
