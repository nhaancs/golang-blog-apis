package productcategorymodel

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	// ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`

	Name string `json:"name" gorm:"column:name;not null;"`
	Slug string `json:"slug" gorm:"column:slug;primaryKey;"`
	Desc string `json:"desc" gorm:"column:desc;"`
}

func (pc ProductCategory) TableName() string {
	return "product_categories"
}

type ProductCategoryCreate struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Slug string `json:"slug"`
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
