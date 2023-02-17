package services

import (
	"errors"

	"github.com/SGDIEGO/ApplicationHA/internal/domain"
	"github.com/SGDIEGO/ApplicationHA/internal/ports"
)

type UserService struct {
	userInterface ports.UserRepositoryI
}

func NewUserService(userInterface ports.UserRepositoryI) ports.UserServiceI {
	return &UserService{
		userInterface: userInterface,
	}
}

func (uS *UserService) GetAllUsers() (*[]domain.User, error) {
	return uS.userInterface.GetUsers()
}

func (uS *UserService) GetUser(id string) (*domain.User, error) {
	return uS.userInterface.GetUserById(id)
}

func (uS *UserService) CreateUser(user domain.User) error {

	if user.Id == "" ||
		user.Name == "" ||
		user.Email == "" ||
		user.Password == "" {
		return errors.New("All required data")
	}

	return uS.userInterface.CreateUser(user)
}

func (uS *UserService) EditUser(id string, newUser domain.User) error {

	return uS.userInterface.EditUser(id, newUser)
}

func (uS *UserService) DeleteUser(id string) error {

	return uS.userInterface.DeleteUser(id)
}
