package domain

import (
	"github.com/google/uuid"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID        string          `json:"id" gorm:"type:char(36);not null;primary_key;uniqueIndex"`
	Name      string          `json:"name" gorm:"char(50); not null"`
	StartDate time.Time       `json:"start_date"`
	EndDate   time.Time       `json:"end_date"`
	CreateAt  *time.Time      `json:"-"`
	UpdatedAt *time.Time      `json:"-"`
	Delete    *gorm.DeletedAt `json:"-"`
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}
