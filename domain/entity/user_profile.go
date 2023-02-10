package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserProfile struct {
	ID        string         `gorm:"size:36;not null;uniqueIndex;primaryKey;" json:"id"`
	UserID    string         `gorm:"size:36;not null;default:null;" json:"user_id"`
	Email     string         `gorm:"size:100" json:"email"`
	Bio       string         `gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

var _ Entity = &UserProfile{}

// BeforeCreate is a hook executed before storing data
func (u *UserProfile) BeforeCreate(tx *gorm.DB) error {
	generateID := uuid.New()
	if u.ID == "" {
		u.ID = generateID.String()
	}

	return nil
}

// TableName is a function return table name
func (u *UserProfile) TableName() string {
	return "user_profiles"
}

// FilterableFields is a function return filterable fields
func (u *UserProfile) FilterableFields() []interface{} {
	return []interface{}{"name"}
}

// TimeFields is a function return time fields
func (u *UserProfile) TimeFields() []interface{} {
	return []interface{}{"created_at", "updated_at", "deleted_at"}
}
