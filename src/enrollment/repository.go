package enrollment

import (
	"log"

	"github.com/NMEJIA93/Api_GO/src/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(enroll *domain.Enrollment) error
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(logger *log.Logger, db *gorm.DB) Repository {
	return &repo{
		db:  db,
		log: logger,
	}
}

func (r *repo) Create(enroll *domain.Enrollment) error {
	if err := r.db.Create(enroll).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}
	r.log.Printf("enrollment created with id: ", enroll.ID)
	return nil
}
