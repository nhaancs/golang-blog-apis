package productmodel

import (
	"errors"
	"strings"
	"time"
)

type ProductUpdate struct {
	UpdatedAt *time.Time `json:"createdAt" gorm:"autoUpdateTime;"`
	Name      string     `json:"name"`
	Desc      *string     `json:"desc"`
}

func (ProductUpdate) TableName() string {
	return Product{}.TableName()
}

func (data *ProductCreate) ValidateUpdate() error {
	data.Name = strings.TrimSpace(data.Name)

	if len(data.Name) == 0 {
		return errors.New("product name can not be blank")
	}

	return nil
}
