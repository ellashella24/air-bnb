package user

import (
	"air-bnb/entities"

	"gorm.io/gorm"
)

type User interface {
	GetAllUser() ([]entities.User, error)
	GetUserByID(userID int) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(userID int, user entities.User) (entities.User, error)
	DeleteUser(userID int) (entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetAllUser() ([]entities.User, error) {
	users := []entities.User{}

	err := ur.db.Find(&users).Error

	if err != nil {
		return users, nil
	}

	return users, nil
}

func (ur *userRepository) GetUserByID(userID int) (entities.User, error) {
	user := entities.User{}

	err := ur.db.Where("ID = ?", userID).Find(&user).Error

	if err != nil {
		return user, nil
	}

	return user, nil
}

func (ur *userRepository) GetUserByEmail(email string) (entities.User, error) {
	user := entities.User{}

	ur.db.Where("email = ?", email).Find(&user)

	return user, nil
}

func (ur *userRepository) CreateUser(newUser entities.User) (entities.User, error) {
	ur.db.Save(&newUser)

	return newUser, nil
}

func (ur *userRepository) UpdateUser(userID int, updatedUser entities.User) (entities.User, error) {
	user := entities.User{}

	err := ur.db.Where("id = ?", userID).Find(&user).Error

	if err != nil {
		return user, err
	}

	ur.db.Model(&user).Updates(updatedUser)

	return updatedUser, nil
}

func (ur *userRepository) DeleteUser(userID int) (entities.User, error) {
	user := entities.User{}

	err := ur.db.Where("id = ?", userID).Delete(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
