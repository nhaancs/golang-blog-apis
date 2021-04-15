package productcategorystore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/productcategory/productcategorymodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *productcategorymodel.ProductCategoryUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}