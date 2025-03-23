package user

import (
	"log"
)

type (
	Service interface {
		Create(dto CreateUserDTO) (*User, error)
		Get(id string) (*User, error)
		GetAll(filter Filters) ([]User, error)
		Delete(id string) error
		Update(id string, firstName *string, lastName *string, email *string, phone *string) error
	}

	service struct {
		log  *log.Logger
		repo Respository
	}

	Filters struct {
		FirstName string
		LastName  string
	}
)

func NewService(log *log.Logger, repo Respository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(dto CreateUserDTO) (*User, error) {
	s.log.Println("Create User Service")
	user := User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Phone:     dto.Phone,
	}
	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s service) GetAll(filter Filters) ([]User, error) {
	users, err := s.repo.GetAll(filter)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s service) Get(id string) (*User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s service) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
	return s.repo.Update(id, firstName, lastName, email, phone)
}
