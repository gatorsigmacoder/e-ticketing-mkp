package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	ADMIN    = "admin"
	USER     = "user"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // UUID as primary key
	Username     string    `json:"username" form:"username" validate:"gte=2,lte=128" gorm:"type:varchar(128);not null;unique"`
	Email        string    `json:"email" form:"email" validate:"required,email" gorm:"type:citext;not null;unique"`
	Password     string    `json:"-" form:"password" validate:"required,gte=8" gorm:"type:varchar(64);not null,colum:password"`
	Role		 string	   `json:"role" form:"role" validate:"gte=2,lte=12" gorm:"type:varchar(12);default:'user'"`
	TokenVersion int       `json:"token_version" form:"token_version" gorm:"default:1"`
	ProfileImage string    `json:"profile_image" form:"profile_image" gorm:"type:text;"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;"`
}