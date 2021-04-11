package productcategorymodel

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID        uint            `json:"id"`
	CreatedAt *time.Time      `json:"createdAt"`
	UpdatedAt *time.Time      `json:"updatedAt,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty"`
	Name      string          `json:"name"`
	Slug      string          `json:"slug"`
	Desc      string          `json:"desc"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}


