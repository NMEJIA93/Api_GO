package user

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Respository interface {
	Create(user *User) error
	GetAll(filters Filters, offset int, limit int) ([]User, error)
	Get(id string) (*User, error)
	Delete(id string) error
	Update(id string, firstName *string, lasName *string, email *string, phone *string) error
	Count(filters Filters) (int, error)
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

func (r *repository) GetAll(filter Filters, offset int, limit int) ([]User, error) {
	var user []User

	tx := r.db.Model(&User{})
	tx = applyFilters(tx, filter)
	tx = tx.Limit(limit).Offset(offset)
	result := tx.Order("Created_at desc").Find(&user)
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

func (r *repository) Count(filters Filters) (int, error) {
	var count int64
	tx := r.db.Model(User{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.FirstName != "" {
		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
		tx = tx.Where("lower(first_name) LIKE ?", filters.FirstName)
	}

	if filters.LastName != "" {
		filters.LastName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.LastName))
		tx = tx.Where("lower(first_name) LIKE ?", filters.LastName)
	}
	return tx
}
