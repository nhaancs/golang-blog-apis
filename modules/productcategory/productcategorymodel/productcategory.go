package productcategorymodel

import (
	"nhaancs/common"
)

const EntityName = "ProductCategory"

type ProductCategory struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Desc            string `json:"desc"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}
