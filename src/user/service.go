package user

import "log"

type Service interface {
	Create(firstName, lastName, email, password string) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s service) Create(firstName, lastName, email, password string) error {
	log.Println("Create User Service")
	return nil
}
