package categorystore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/categorymodel"
	"time"
)

func (s *sqlStore) SoftDelete(
	ctx context.Context,
	id int,
) error {
	db := s.db
	err := db.Table(categorymodel.Category{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
