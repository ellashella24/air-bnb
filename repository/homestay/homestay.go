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
	GetAllHostHomestay(id int) ([]entities.Homestay, error)

	// CREATE
	CreteaHomestay(homestay entities.Homestay) (entities.Homestay, error)
	// UPDATE
	GetHomestayIdByHostId(id int) (entities.Homestay, error)
	UpdateHomestay(homestay entities.Homestay) (entities.Homestay, error)
	// DELETE
	DeleteHomestayByHostId(id int) error
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

func (h *homestay) GetAllHostHomestay(id int) ([]entities.Homestay, error) {
	var homestays []entities.Homestay
	err := h.db.Where("host_id = ?", id).Find(&homestays).Error
	if err != nil {
		return nil, err
	}
	return homestays, nil
}

func (h *homestay) CreteaHomestay(homestay entities.Homestay) (entities.Homestay, error) {

	err := h.db.Save(&homestay).Error
	if err != nil {
		return homestay, err
	}
	return homestay, err
}

func (h *homestay) GetHomestayIdByHostId(id int) (entities.Homestay, error) {
	var homestay entities.Homestay
	// err := h.db.First(&homestay, id).Error
	err := h.db.Where("host_id = ?", id).Find(&homestay).Error
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

func (h *homestay) DeleteHomestayByHostId(id int) error {
	var delete entities.Homestay
	// err := h.db.Delete(&delete).Where("host_id = ?", id).Error
	err := h.db.Where("host_id = ?", id).Delete(&delete).Error
	if err != nil {
		return err
	}
	return nil
}
