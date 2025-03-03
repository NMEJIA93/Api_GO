package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string          `json:"id" gorm:"char(36); not null; primary_key; unique_index"`
	FirstName string          `json:"first_name" gorm:"char(50); not null"`
	LastName  string          `json:"last_name" gorm:"char(50); not null"`
	Email     string          `json:"email" gorm:"char(50); not null; unique_index"`
	Phone     string          `json:"phone" gorm:"char(30); not null; unique_index"`
	CreatedAt *time.Time      `json:"-"`
	UpdatedAt *time.Time      `json:"-"`
	Delete    *gorm.DeletedAt `json:"-"`
}
