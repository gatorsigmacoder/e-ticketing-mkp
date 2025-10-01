package models

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // UUID as primary key
	UserId       uuid.UUID `json:"user_id" form:"user_id" validate:"required" gorm:"type:uuid;not null"`
	Balance      float64   `json:"balance" form:"balance" validate:"required" gorm:"type:double precision;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;"`
}