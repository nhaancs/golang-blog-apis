package categorystore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/model"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *categorymodel.CategoryUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
