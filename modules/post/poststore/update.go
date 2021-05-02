package poststore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/postmodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *postmodel.PostUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
