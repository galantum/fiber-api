package services

import (
	"fiber-api/models"
	"fiber-api/repositories"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	GetUser(id int) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(user models.User) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.userRepo.GetUsers()
}

func (s *userService) GetUser(id int) (models.User, error) {
	return s.userRepo.GetUser(id)
}

func (s *userService) CreateUser(user models.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *userService) UpdateUser(user models.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *userService) DeleteUser(user models.User) error {
	return s.userRepo.DeleteUser(user)
}
