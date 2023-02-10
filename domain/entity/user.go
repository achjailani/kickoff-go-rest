package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is struct defines user entity
type User struct {
	ID          string         `gorm:"size:36;not null;uniqueIndex;primaryKey;" json:"id"`
	RoleID      string         `gorm:"size:36;not null;default:null;" json:"role_id"`
	Role        Role           `gorm:"foreignKey:RoleID;references:ID"`
	Name        string         `gorm:"size:100;not null;" json:"name"`
	Username    string         `gorm:"size:100;not null;uniqueIndex;" json:"username"`
	Password    string         `gorm:"size:255;not null;" json:"password"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	UserProfile UserProfile    `json:"user_profile"`
}

// Implements base entity methods
var _ Entity = &User{}

// BeforeCreate is a hook executed before storing data
func (u *User) BeforeCreate(tx *gorm.DB) error {
	generateID := uuid.New()
	if u.ID == "" {
		u.ID = generateID.String()
	}

	return nil
}

// TableName is a function return table name
func (u *User) TableName() string {
	return "users"
}

// FilterableFields is a function return filterable fields
func (u *User) FilterableFields() []interface{} {
	return []interface{}{"name"}
}

// TimeFields is a function return time fields
func (u *User) TimeFields() []interface{} {
	return []interface{}{"created_at", "updated_at", "deleted_at"}
}
