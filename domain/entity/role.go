package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Role is struct defines role entity
type Role struct {
	ID          string         `gorm:"size:36;not null;uniqueIndex;primaryKey;" json:"id"`
	Name        string         `gorm:"size:100;not null;" json:"name"`
	Code        string         `gorm:"size:100;not null;uniqueIndex;" json:"code"`
	Description string         `gorm:"size:255;" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

// Implements base entity methods
var _ Entity = &Role{}

// BeforeCreate is a hook executed before creating data
func (u *Role) BeforeCreate(_ *gorm.DB) error {
	generateID := uuid.New()
	if u.ID == "" {
		u.ID = generateID.String()
	}

	return nil
}

// TableName is a function return table name
func (u *Role) TableName() string {
	return "roles"
}

// FilterableFields is a function return filterable fields
func (u *Role) FilterableFields() []interface{} {
	return []interface{}{"name", "code", "id"}
}

// TimeFields is a function return time fields
func (u *Role) TimeFields() []interface{} {
	return []interface{}{"created_at", "updated_at", "deleted_at"}
}
