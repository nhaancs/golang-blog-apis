package productmodel

import (
	"nhaancs/common"
)

const EntityName = "Product"
const TableName = "products"

type Product struct {
	common.SQLModel `json:",inline"`
	Name            string  `json:"name"`
	Slug            string  `json:"slug"`
	ShortDesc       string  `json:"shortDesc" db:"short_desc"`
	LongDesc        string  `json:"longDesc" db:"long_desc"`
	UnitKey         string  `json:"unitKey" db:"unit_key"`
	UnitName        string  `json:"unitName" db:"unit_name"`
	Price           float64 `json:"price"`
	Quantity        float64 `json:"quantity"`
	IsUnlimited     bool    `json:"isUnlimited" db:"is_unlimited"`
	IsEnabled       bool    `json:"isEnabled" db:"is_enabled"`
	//todo: images
}
