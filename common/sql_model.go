package common

import "time"

type SQLModel struct {
	Id        uint       `json:"id"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
