package categorystore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/model"
)

func (s *sqlStore) Create(ctx context.Context, data *categorymodel.CategoryCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
