package models

import (
	"time"

	"github.com/google/uuid"
)

type Terminal struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // UUID as primary key
	Name     	 string    `json:"name" form:"name" validate:"gte=2,lte=128" gorm:"type:varchar(128);not null;unique"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;"`
}