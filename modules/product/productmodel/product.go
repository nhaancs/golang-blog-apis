package productmodel

import (
	"nhaancs/common"
)

const EntityName = "Product"

type Product struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Desc            string `json:"desc"`
}

func (Product) TableName() string {
	return "product_categories"
}
