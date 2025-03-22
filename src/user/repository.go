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
	Update(id string, firstName *string, lasName *string, email *string, phone *string) error
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

	err := r.db.First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Delete(id string) error {
	user := User{ID: id}
	//Eliminado Fisico
	//result := r.db.Delete(&user)
	err := r.db.First(&user).Error
	if err != nil {
		return err
	}

	result := r.db.Delete(&user).Error

	if result != nil {
		return result
	}
	
	return nil
}

func (r *repository) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
	values := make(map[string]interface{})

	if firstName != nil {
		values["first_name"] = *firstName
	}
	if lastName != nil {
		values["last_name"] = *lastName
	}
	if email != nil {
		values["email"] = *email
	}
	if phone != nil {
		values["phone"] = *phone
	}

	if err := r.db.Model(&User{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil
}
