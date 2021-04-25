package productmodel

import (
	"errors"
	"strings"
)

type ProductUpdate struct {
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	ShortDesc   string  `json:"shortDesc" db:"short_desc"`
	LongDesc    *string `json:"longDesc" db:"long_desc"`
	UnitKey     string  `json:"unitKey" db:"unit_key"`
	UnitName    string  `json:"unitName" db:"unit_name"`
	Price       float32 `json:"price"`
	Quantity    float32 `json:"quantity"`
	IsUnlimited bool    `json:"isUnlimited" db:"is_unlimited"`
	IsEnabled   bool    `json:"isEnabled" db:"is_enabled"`
}

func (data *ProductUpdate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	data.ShortDesc = strings.TrimSpace(data.ShortDesc)
	data.UnitKey = strings.TrimSpace(data.UnitKey)

	if len(data.Name) == 0 {
		return errors.New("product name can not be blank")
	}
	if len(data.Name) > 255 {
		return errors.New("product name is too long")
	}

	if len(data.ShortDesc) == 0 {
		return errors.New("short description can not be blank")
	}
	if len(data.ShortDesc) > 255 {
		return errors.New("short description is too long")
	}

	if len(data.UnitKey) == 0 {
		return errors.New("unit can not be blank")
	}
	if len(data.UnitKey) > 50 {
		return errors.New("invalid unit")
	}

	if data.Price < 0 || data.Price > 9999999999999.99 {
		return errors.New("invalid price")
	}

	if data.Quantity < 0 || data.Quantity > 9999999999999.99 {
		return errors.New("invalid quantity")
	}

	return nil
}
