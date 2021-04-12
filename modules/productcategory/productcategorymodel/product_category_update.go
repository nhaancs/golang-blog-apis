package productcategorymodel

import (
	"errors"
	"strings"
	"time"
)

type ProductCategoryUpdate struct {
	UpdatedAt *time.Time `json:"createdAt"`
	Name      string     `json:"name"`
	Desc      *string     `json:"desc"`
}

func (ProductCategoryUpdate) TableName() string {
	return ProductCategory{}.TableName()
}

func (data *ProductCategoryCreate) ValidateUpdate() error {
	data.Name = strings.TrimSpace(data.Name)

	if len(data.Name) == 0 {
		return errors.New("product category name can not be blank")
	}

	return nil
}
