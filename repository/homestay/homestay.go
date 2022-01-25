package homestay

import (
	"air-bnb/entities"

	"gorm.io/gorm"
)

//--------------------------------------------------------------------------------
//INTERFACE
//--------------------------------------------------------------------------------
type InterfaceHomestay interface {
	GetallHomestay() ([]entities.Homestay, error)
	GetHostHomestay(id int) ([]entities.Homestay, error)

	//CREATE, UPDATE, DELETE
	CreteaHomestay(homestay entities.Homestay) (entities.Homestay, error)
	GetHomestayId(id int) (entities.Homestay, error)
	UpdateHomestay(homestay entities.Homestay) (entities.Homestay, error)
	DeleteHomestay(id int) error
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
		return nil, err
	}
	return homestays, nil
}

func (h *homestay) GetHostHomestay(id int) ([]entities.Homestay, error) {
	var homestays []entities.Homestay
	err := h.db.Where("host_id = ?", id).Find(&homestays).Error
	if err != nil {
		return nil, err
	}
	return homestays, nil
}

func (h *homestay) GetHomestayId(id int) (entities.Homestay, error) {
	var homestay entities.Homestay
	err := h.db.First(&homestay, id).Error
	if err != nil {
		return homestay, err
	}
	return homestay, err
}

func (h *homestay) UpdateHomestay(homestay entities.Homestay) (entities.Homestay, error) {
	err := h.db.Save(&homestay).Error
	if err != nil {
		return homestay, err
	}
	return homestay, err
}

func (h *homestay) CreteaHomestay(homestay entities.Homestay) (entities.Homestay, error) {
	err := h.db.Save(&homestay).Error
	if err != nil {
		return homestay, err
	}
	return homestay, err
}

func (h *homestay) DeleteHomestay(id int) error {
	var delete entities.Homestay
	err := h.db.Delete(&delete, id).Error
	if err != nil {
		return err
	}
	return nil
}
