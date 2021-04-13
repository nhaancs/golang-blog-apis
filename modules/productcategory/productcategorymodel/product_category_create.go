package productcategorymodel

import (
	"errors"
	"strings"
	"time"
)

type ProductCategoryCreate struct {
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime;"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Slug      string     `json:"slug"`
}

func (ProductCategoryCreate) TableName() string {
	return ProductCategory{}.TableName()
}

func (data *ProductCategoryCreate) ValidateCreate() error {
	data.Name = strings.TrimSpace(data.Name)

	if len(data.Name) == 0 {
		return errors.New("product category name can not be blank")
	}

	return nil
}
