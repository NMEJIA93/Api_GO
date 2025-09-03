package enrollment

import (
	"errors"
	"log"

	"github.com/NMEJIA93/Api_GO/src/course"
	"github.com/NMEJIA93/Api_GO/src/domain"
	"github.com/NMEJIA93/Api_GO/src/user"
)

type (
	Service interface {
		//Create(userID, courseID string) (*domain.Enrollment, error)
		Create(enrollDto CreateEnrollmentDTO) (*domain.Enrollment, error)
	}
	service struct {
		log       *log.Logger
		userSrv   user.Service
		courseSrv course.Service
		repo      Repository
	}
)

func NewService(l *log.Logger, userSrv user.Service, courseSrv course.Service, repo Repository) Service {
	return &service{
		log:       l,
		userSrv:   userSrv,
		courseSrv: courseSrv,
		repo:      repo,
	}
}

func (s service) Create(enrollDto CreateEnrollmentDTO) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID:   enrollDto.UserID,
		CourseID: enrollDto.CourseID,
		Status:   "P", // pendiente
	}

	if _, err := s.userSrv.Get(enroll.UserID); err != nil {
		s.log.Printf("error getting (User doesn´t exist) user: %v", err)
		return nil, errors.New("user does not exist")
	}

	if _, err := s.courseSrv.Get(enroll.CourseID); err != nil {
		s.log.Printf("error getting (Course doesn´t exist) course: %v", err)
		return nil, errors.New("course does not exist")
	}

	if err := s.repo.Create(enroll); err != nil {
		s.log.Printf(" error %v", err)
		return nil, err
	}

	return enroll, nil
}
