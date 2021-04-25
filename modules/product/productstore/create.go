package productstore

import (
	"context"
	"fmt"
	"nhaancs/common"
	"nhaancs/modules/product/productmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *productmodel.ProductCreate) (lastInsertId int64, err error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (name, slug, short_desc, long_desc, unit_key, unit_name, price, quantity, is_unlimited, is_enabled) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 
		productmodel.TableName,
	)
	res, err := s.db.Exec(query, data.Name, data.Slug, data.ShortDesc, data.LongDesc, data.UnitKey, data.UnitName, data.Price, data.Quantity, data.IsUnlimited, data.IsEnabled)
	if err != nil {
		return 0, common.ErrDB(err)
	}

	return res.LastInsertId()
}