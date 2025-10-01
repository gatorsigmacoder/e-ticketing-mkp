package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           	uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // UUID as primary key
	TripId       	uuid.UUID `json:"trip_id" form:"trip_id" validate:"required" gorm:"type:uuid;not null"`
	TotalFare       float64   `json:"total_fare" form:"total_fare" validate:"required" gorm:"type:double precision;not null"`
	CreatedAt    	time.Time `json:"created_at" gorm:"type:timestamp;"`
	UpdatedAt    	time.Time `json:"updated_at" gorm:"type:timestamp;"`
}