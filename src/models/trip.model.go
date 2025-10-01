package models

import (
	"time"

	"github.com/google/uuid"
)

type Trip struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // UUID as primary key
	UserId       uuid.UUID `json:"user_id" form:"user_id" validate:"required" gorm:"type:uuid;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;"`
}