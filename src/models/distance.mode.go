package models

import (
	"time"

	"github.com/google/uuid"
)

type Distance struct {
	ID           	uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // UUID as primary key
	FromTerminal    uuid.UUID `json:"from_terminal" form:"from_terminal" validate:"required" gorm:"type:uuid;not null"`
	ToTerminal      uuid.UUID `json:"to_terminal" form:"to_terminal" validate:"required" gorm:"type:uuid;not null"`
	Fare            float64   `json:"fare" form:"fare" validate:"required" gorm:"type:double precision;not null"`
	CreatedAt    	time.Time `json:"created_at" gorm:"type:timestamp;"`
	UpdatedAt    	time.Time `json:"updated_at" gorm:"type:timestamp;"`
}