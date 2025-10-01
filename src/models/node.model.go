package models

import (
	"time"

	"github.com/google/uuid"
)

type Node struct {
	ID           	uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // UUID as primary key
	TripId       	uuid.UUID `json:"trip_id" form:"trip_id" validate:"required" gorm:"type:uuid;not null"`
	TerminalId   	uuid.UUID `json:"terminal_id" form:"terminal_id" validate:"required" gorm:"type:uuid;not null"`
	FirstTerminal   bool      `json:"first_terminal" form:"first_terminal" gorm:"type:boolean;default:false"`
	LastTerminal    bool      `json:"last_terminal" form:"last_terminal" gorm:"type:boolean;default:false"`
	CreatedAt    	time.Time `json:"created_at" gorm:"type:timestamp;"`
	UpdatedAt    	time.Time `json:"updated_at" gorm:"type:timestamp;"`
}