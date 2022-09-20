package services

import (
	"agmc_d6/models"
	"agmc_d6/repositories"
)

type UserService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewUser(userRepo repositories.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) FindAll() ([]models.User, error) {
	users, err := us.userRepo.FindAll()
	return users, err
}

func (us *UserService) Find(id int) (models.User, error) {
	user, err := us.userRepo.Find(id)
	return user, err
}

func (us *UserService) Insert(user models.User) (models.User, error) {
	user, err := us.userRepo.Create(user)
	return user, err
}

func (us *UserService) Update(id int, user models.User) (models.User, error) {
	user, err := us.userRepo.Update(user, id)
	return user, err
}

func (us *UserService) Delete(id int) (error) {
	err := us.userRepo.Delete(id)
	return err
}