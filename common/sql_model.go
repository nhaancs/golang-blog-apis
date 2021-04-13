package common

import "time"

type SQLModel struct {
	Id        uint       `json:"id" gorm:"primaryKey;autoIncrement;"`
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"autoUpdateTime;"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index;"`
}
