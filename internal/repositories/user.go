package repository

import (
	"errors"

	"github.com/SGDIEGO/ApplicationHA/internal/domain"
	"github.com/SGDIEGO/ApplicationHA/internal/ports"
)

type UserRepository struct {
	users []domain.User
}

func NewMemKVS(data []domain.User) ports.UserRepositoryI {
	return &UserRepository{users: data}
}

func (uR *UserRepository) GetUsers() (*[]domain.User, error) {

	return &uR.users, nil
}

func (uR *UserRepository) GetUserById(id string) (*domain.User, error) {

	for _, u := range uR.users {

		if u.Id == id {
			return &u, nil
		}
	}

	return &domain.User{}, errors.New("game not found in kvs")
}

func (uR *UserRepository) CreateUser(user domain.User) error {

	// Add new user
	uR.users = append(uR.users, user)

	return nil
}

func (uR *UserRepository) EditUser(id string, newUser domain.User) error {

	for i, user := range uR.users {
		if user.Id == id {
			uR.users[i] = newUser
			return nil
		}
	}

	return errors.New("not found user")
}

func (uR *UserRepository) DeleteUser(id string) error {

	for i, user := range uR.users {
		if user.Id == id {
			ret := make([]domain.User, 0)
			ret = append(ret, uR.users[:i]...)
			ret = append(ret, uR.users[i+1:]...)

			uR.users = ret
			return nil
		}
	}

	return errors.New("not found user")
}
