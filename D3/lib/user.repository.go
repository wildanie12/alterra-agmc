package lib

import (
	"agmc_d3/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) FindAll() ([]models.User, error) {
	users := []models.User{}
	tx := ur.db.Find(&users)
	if tx.Error != nil {
		return []models.User{}, errors.New(fmt.Sprintf("cannot find all users: %v", tx.Error))
	}
	return users, nil
}

func (ur *UserRepository) Find(id int) (models.User, error) {
	user := models.User{}
	tx := ur.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return models.User{}, errors.New(fmt.Sprintf("cannot get user detail: %v", tx.Error))
	}
	return user, nil
}

func (ur *UserRepository) Create(user models.User) (models.User, error) {
	tx := ur.db.Create(&user)
	if tx.Error != nil {
		return models.User{}, errors.New(fmt.Sprintf("cannot get user detail: %v", tx.Error))
	}
	return user, nil
}

func (ur *UserRepository) Update(user models.User, id int) (models.User, error) {
	tx := ur.db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return models.User{}, errors.New(fmt.Sprintf("cannot update user: %v", tx.Error))
	} else if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New(fmt.Sprintf("cannot get user detail: %v", tx.Error))
	}
	return user, nil
}

func (ur *UserRepository) Delete(id int) error {
	tx := ur.db.Where("id = ?", id).Delete(&models.User{})
	if tx.Error != nil {
		return errors.New(fmt.Sprintf("cannot delete user: %v", tx.Error))
	} else if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("cannot get user detail: %v", tx.Error))
	}
	return nil
}