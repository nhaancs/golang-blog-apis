package poststore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/model"
)

func (s *sqlStore) Create(ctx context.Context, data *postmodel.PostCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
