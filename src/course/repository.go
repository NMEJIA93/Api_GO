package course

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"strings"
)

type (
	Repository interface {
		Create(course *Course) error
		GetByID(id string) (*Course, error)
		GetAll(filter Filters, offset int, limit int) ([]Course, error)
		Count(filters Filters) (int, error)
		Delete(id string) error
	}

	repository struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepository(log *log.Logger, db *gorm.DB) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}

func (r *repository) Create(course *Course) error {
	if err := r.db.Create(course).Error; err != nil {
		r.log.Printf("Error while creating course: %v", err)
		return err
	}
	r.log.Println("Course created with id: ", course.ID)
	return nil
}

func (r *repository) GetByID(id string) (*Course, error) {
	course := Course{ID: id}
	err := r.db.First(&course).Error
	if err != nil {
		r.log.Printf("Error while getting course by id: %v", err)
		return nil, err
	}
	return &course, nil
}

func (r *repository) GetAll(filter Filters, offset int, limit int) ([]Course, error) {
	var courses []Course

	tx := r.db.Model(&Course{})
	tx = applyFilters(tx, filter)
	err := r.db.Find(&courses).Error
	if err != nil {
		r.log.Printf("Error while getting all courses: %v", err)
		return nil, err
	}
	return courses, nil
}

func (r *repository) Count(filters Filters) (int, error) {
	var count int64
	tx := r.db.Model(Course{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *repository) Delete(id string) error {
	course := Course{ID: id}
	err := r.db.First(&course).Error
	if err != nil {
		return err
	}

	result := r.db.Delete(&course).Error
	if result != nil {
		return result
	}
	return nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("lower(first_name) LIKE ?", filters.Name)
	}
	return tx
}
