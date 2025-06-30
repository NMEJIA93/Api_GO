package course

import (
	"log"
	"time"
)

type (
	Service interface {
		Create(dto CreateCourseDTO) (*Course, error)
		Get(id string) (*Course, error)
		GetAll(filter Filters, offset int, limit int) ([]Course, error)
		Count(filter Filters) (int, error)
		Delete(id string) error
		Update(course UpdateCourseDTO) error
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

func (s service) Update(course UpdateCourseDTO) error {

	var startDateParsed, endDateParsed *time.Time
	if &course.StartDate != nil {
		date, err := time.Parse("2006-01-02", *course.StartDate)
		if err != nil {
			return err
		}
		startDateParsed = &date
	}
	if &course.EndDate != nil {
		date, err := time.Parse("2006-01-02", *course.EndDate)
		if err != nil {
			return err
		}
		endDateParsed = &date
	}

	return s.repo.Update(course.ID, course.Name, startDateParsed, endDateParsed)
}

func (s service) Get(id string) (*Course, error) {
	s.log.Println("Get Course Service")
	course, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return course, nil
}
func (s service) GetAll(filter Filters, offset int, limit int) ([]Course, error) {
	s.log.Println("GetAll Course Service")
	courses, err := s.repo.GetAll(filter, offset, limit)
	if err != nil {
		return nil, err
	}
	return courses, nil
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

func (s service) Count(filter Filters) (int, error) {
	return s.repo.Count(filter)
}

func (s service) Delete(id string) error {
	return s.repo.Delete(id)
}
