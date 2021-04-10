package productcategorymodel

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	Desc      string         `json:"desc"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}

type ProductCategoryCreate struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Slug      string    `json:"slug"`
}

func (ProductCategoryCreate) TableName() string {
	return ProductCategory{}.TableName()
}

func (data *ProductCategoryCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if len(data.Name) == 0 {
		return errors.New("product category name can not be blank")
	}

	return nil
}
