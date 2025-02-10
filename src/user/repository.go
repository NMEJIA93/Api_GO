package user

import (
	"gorm.io/gorm"
	"log"

	"github.com/google/uuid"
)

type Respository interface {
	Create(user *User) error
}

type repository struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepository(log *log.Logger, db *gorm.DB) Respository {
	return &repository{
		log: log,
		db:  db,
	}
}

func (r *repository) Create(user *User) error {
	user.ID = uuid.New().String()

	result := r.db.Create(user)
	if result.Error != nil {
		r.log.Printf("Error while creating user: %v", result.Error)
		return result.Error
	}

	r.log.Println("user Created with id: ", user.ID)
	return nil
}
