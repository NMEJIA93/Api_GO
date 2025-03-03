package user

import (
	"gorm.io/gorm"
	"log"

	"github.com/google/uuid"
)

type Respository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	Get(id string) (*User, error)
	Delete(id string) error
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

	if err := r.db.Create(user).Error; err != nil {
		r.log.Printf("Error while creating user: %v", err)
		return err
	}

	r.log.Println("user Created with id: ", user.ID)
	return nil
}

func (r *repository) GetAll() ([]User, error) {
	var user []User
	result := r.db.Model(&user).Order("created_at desc").Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *repository) Get(id string) (*User, error) {
	user := User{ID: id}

	result := r.db.First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repository) Delete(id string) error {
	user := User{ID: id}
	//Eliminado Fisico
	//result := r.db.Delete(&user)
	result := r.db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
