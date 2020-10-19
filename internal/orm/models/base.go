package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// BaseModel defines common columns in all DB structs
type BaseModel struct {
	ID        uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"index"`
}

// BaseModelSoftDelete defines common columns that allow soft delete
type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `sql:"index"`
}
