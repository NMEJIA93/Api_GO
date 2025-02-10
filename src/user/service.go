package user

import "log"

type Service interface {
	Create(dto CreateUserDTO) error
}

type service struct {
	log  *log.Logger
	repo Respository
}

func NewService(log *log.Logger, repo Respository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(dto CreateUserDTO) error {
	s.log.Println("Create User Service")
	user := User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Phone:     dto.Phone,
	}
	s.repo.Create(&user)
	return nil
}
