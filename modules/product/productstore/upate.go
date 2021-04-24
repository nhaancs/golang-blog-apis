package productstore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/product/productmodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *productmodel.ProductUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}