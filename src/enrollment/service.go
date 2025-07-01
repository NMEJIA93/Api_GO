package enrollment

import (
	"log"

	"github.com/NMEJIA93/Api_GO/src/domain"
)

type (
	Service interface {
		//Create(userID, courseID string) (*domain.Enrollment, error)
		Create(enrollDto CreateEnrollmentDTO) (*domain.Enrollment, error)
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) Create(enrollDto CreateEnrollmentDTO) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID:   enrollDto.UserID,
		CourseID: enrollDto.CourseID,
		Status:   "P",
	}

	if err := s.repo.Create(enroll); err != nil {
		s.log.Printf(" error %v", err)
		return nil, err
	}

	return enroll, nil
}
