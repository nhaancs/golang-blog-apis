package productcategorystore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/productcategory/productcategorymodel"
	"time"
)

// todo: test list after delete

func (s *sqlStore) SoftDelete(
	ctx context.Context,
	id int,
) error {
	db := s.db
	if err := db.Table(productcategorymodel.ProductCategory{}.TableName()).Where("id = ?", id).Updates(
		map[string]interface{}{"deleted_at": time.Now()}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}