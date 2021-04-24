package productmodel

import (
	"errors"
	"strings"
	"time"
)

type ProductCreate struct {
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime;"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Slug      string     `json:"slug"`
}

func (ProductCreate) TableName() string {
	return Product{}.TableName()
}

func (data *ProductCreate) ValidateCreate() error {
	data.Name = strings.TrimSpace(data.Name)

	if len(data.Name) == 0 {
		return errors.New("product name can not be blank")
	}

	return nil
}
