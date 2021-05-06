package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	IsEnabled bool       `json:"is_enabled" gorm:"column:is_enabled;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at;"`
}
type SQLCreateModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	IsEnabled bool       `json:"is_enabled" gorm:"column:is_enabled;default:true;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime;"`
}
type SQLUpdateModel struct {
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;"`
	IsEnabled *bool       `json:"is_enabled,omitempty" gorm:"column:is_enabled;"`
}

func (m *SQLModel) GenUID(dbType int) {
	m.FakeId = NewUID(uint32(m.Id), dbType, 1)
}

func (m *SQLCreateModel) GenUID(dbType int) {
	m.FakeId = NewUID(uint32(m.Id), dbType, 1)
}
