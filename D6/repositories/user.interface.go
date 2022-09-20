package repositories

import "agmc_d6/models"

type UserRepositoryInterface interface {
	// FindAll - .
	FindAll() ([]models.User, error) 

	// FindByUsername - .
	FindByUsername(username string) (models.User, error)

	// Find - .
	Find(id int) (models.User, error) 

	// Create - .
	Create(user models.User) (models.User, error) 

	// Update - .
	Update(user models.User, id int) (models.User, error) 

	// Delete - .
	Delete(id int) error
}