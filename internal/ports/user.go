package ports

import "github.com/SGDIEGO/ApplicationHA/internal/domain"

type UserRepositoryI interface {
	GetUsers() (*[]domain.User, error)
	GetUserById(id string) (*domain.User, error)
	CreateUser(user domain.User) error
	EditUser(id string, newUser domain.User) error
	DeleteUser(id string) error
}

type UserServiceI interface {
	GetAllUsers() (*[]domain.User, error)
	GetUser(id string) (*domain.User, error)
	CreateUser(user domain.User) error
	EditUser(id string, newUser domain.User) error
	DeleteUser(id string) error
}
