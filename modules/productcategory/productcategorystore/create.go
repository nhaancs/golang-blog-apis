package productcategorystore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/productcategory/productcategorymodel"
)

func (s *sqlStore) Create(ctx context.Context, data *productcategorymodel.ProductCategoryCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}