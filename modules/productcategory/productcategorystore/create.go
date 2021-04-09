package productcategorystore

import (
	"context"
	"nhaancs/modules/productcategory/productcategorymodel"
)

func (s *sqlStore) Create(ctx context.Context, data *productcategorymodel.ProductCategoryCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}