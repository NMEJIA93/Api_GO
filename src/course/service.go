package course

import (
	"log"
	"time"
)

type (
	Service interface {
		Create(dto CreateCourseDTO) (*Course, error)
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
	Filters struct {
		Name string
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}
func (s service) Create(dto CreateCourseDTO) (*Course, error) {
	s.log.Println("Create Course Service")

	startDateParsed, err := time.Parse("2006-01-02", dto.StartDate)
	if err != nil {
		s.log.Printf("Error while parsing start date: %v", err)
		return nil, err
	}
	endDateParsed, err := time.Parse("2006-01-02", dto.StartDate)
	if err != nil {
		s.log.Printf("Error while parsing start date: %v", err)
		return nil, err
	}

	course := &Course{
		Name:      dto.Name,
		StartDate: startDateParsed,
		EndDate:   endDateParsed,
	}
	if err := s.repo.Create(course); err != nil {
		return nil, err
	}
	return course, nil
}
