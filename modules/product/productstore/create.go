package productstore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/product/productmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *productmodel.ProductCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}